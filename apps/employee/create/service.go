package create

import (
	"context"
	"database/sql"
	"log/slog"
)

type repositoryCreateEmployeeContract interface {
	FindAuthByEmail(ctx context.Context, email string) (auth Auth, err error)
	CreateAuth(ctx context.Context, tx *sql.Tx, auth Auth) error

	CreateEmployee(ctx context.Context, tx *sql.Tx, employee Employee) error

	// db transaction
	Begin(ctx context.Context) (db *sql.Tx, err error)
	Commit(ctx context.Context, tx *sql.Tx) error
	Rollback(ctx context.Context, tx *sql.Tx) error
}

type createEmployeeService struct {
	repository repositoryCreateEmployeeContract
}

func NewCreateEmployeeService(repository repositoryCreateEmployeeContract) createEmployeeService {
	return createEmployeeService{repository: repository}
}

func (s createEmployeeService) create(ctx context.Context, req CreateEmployeeRequest) (err error) {
	// check auth is exist to db
	auth, err := s.repository.FindAuthByEmail(ctx, req.Email)

	if err != nil {
		if err != errEmailNotFound {
			slog.ErrorContext(ctx, "[create] error when FindAuthByEmail", slog.String("error", err.Error()))
			return err
		}
	}

	if auth.IsExist() {
		slog.ErrorContext(ctx, "[create] email is exist")
		return errEmailAlreadyRegistered
	}

	// start db transaction
	tx, err := s.repository.Begin(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "[create] error when start db transaction", slog.Any("error", err.Error()))
		return err
	}

	// if error/failed -> rollback
	defer s.repository.Rollback(ctx, tx)

	// continue db transaction
	authModel := req.ToAuthModel()
	if err := s.repository.CreateAuth(ctx, tx, authModel); err != nil {
		slog.ErrorContext(ctx, "[create] error when try to CreateAuth", slog.Any("error", err.Error()))
		return err
	}

	employeeModel := req.ToEmployeeModel(authModel.PublicId)
	if err := s.repository.CreateEmployee(ctx, tx, employeeModel); err != nil {
		slog.ErrorContext(ctx, "[create] error when try to CreateEmployee", slog.Any("error", err.Error()))
		return err
	}

	// commit db transaction
	err = s.repository.Commit(ctx, tx)
	if err != nil {
		slog.ErrorContext(ctx, "[create] error when try to commit db transaction", slog.Any("error", err.Error()))
		return err
	}

	return
}
