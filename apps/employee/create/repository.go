package create

import (
	"context"
	"database/sql"
)

type createEmployeeRepository struct {
	db *sql.DB
}

// Begin implements [repositoryCreateEmployeeContract].
func (c createEmployeeRepository) Begin(ctx context.Context) (db *sql.Tx, err error) {
	return c.db.BeginTx(ctx, &sql.TxOptions{})
}

// Commit implements [repositoryCreateEmployeeContract].
func (c createEmployeeRepository) Commit(ctx context.Context, tx *sql.Tx) error {
	return tx.Commit()
}

// CreateAuth implements [repositoryCreateEmployeeContract].
func (c createEmployeeRepository) CreateAuth(ctx context.Context, tx *sql.Tx, auth Auth) error {
	query := `
		INSERT INTO nb_food_app.auth (
			public_id
			, email
			, password
			, role
			, is_active
			, created_at
			, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, now(), now() 
		)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		auth.PublicId,
		auth.Email,
		auth.Password,
		auth.Role,
		auth.IsActive,
	)

	return err
}

// CreateEmployee implements [repositoryCreateEmployeeContract].
func (c createEmployeeRepository) CreateEmployee(ctx context.Context, tx *sql.Tx, employee Employee) error {
	query := `
		INSERT INTO nb_food_app.employees (
			public_id
			, name
			, profile
			, auth_id
			, created_at
			, updated_at
		) VALUES (
			$1, $2, $3, $4, now(), now() 
		)
	`
	_, err := tx.ExecContext(
		ctx,
		query,
		employee.PublicId,
		employee.Name,
		employee.Profile,
		employee.AuthId,
	)

	return err
}

// FindAuthByEmail implements [repositoryCreateEmployeeContract].
func (c createEmployeeRepository) FindAuthByEmail(ctx context.Context, email string) (auth Auth, err error) {
	query := `
		SELECT
			email
		FROM auth
		WHERE email = $q
	`

	row := c.db.QueryRowContext(ctx, query, email)
	err = row.Scan(&auth.Email)

	if err == sql.ErrNoRows {
		return Auth{}, errEmailNotFound
	} else if err != nil {
		return Auth{}, nil
	}

	return
}

// Rollback implements [repositoryCreateEmployeeContract].
func (c createEmployeeRepository) Rollback(ctx context.Context, tx *sql.Tx) error {
	return tx.Rollback()
}

func NewCreateEmployeeRepository(db *sql.DB) createEmployeeRepository {
	return createEmployeeRepository{db: db}
}
