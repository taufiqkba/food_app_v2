package login

import (
	"context"
	"log/slog"
	"time"

	"github.com/taufiqkba/food_app_v2/internal/config"
	token "github.com/taufiqkba/food_app_v2/internal/utils/jwt"
)

type repoContract interface {
	GetByEmail(ctx context.Context, email string) (auth Auth, err error)
}

type serviceLogin struct {
	repo repoContract
}

func newServiceLogin(repo repoContract) serviceLogin {
	return serviceLogin{repo: repo}
}

func (s serviceLogin) login(ctx context.Context, req LoginRequest) (tok string, role string, err error) {
	// check to db
	auth, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		slog.ErrorContext(ctx, "[lgin] error when try to GetByEmail", slog.Any("error", err.Error()))
	}

	// check status user
	if !auth.IsActive {
		slog.ErrorContext(ctx, "[login] user is not active", slog.Any("error", err.Error()))
		return "", "", nil
	}

	// check password
	if err = auth.ValidatePassowrd(req.Password); err != nil {
		slog.ErrorContext(ctx, "[login] error when validate password", slog.Any("error", err.Error()))
		return "", "", err
	}

	// create token
	cfg := config.GetConfig()
	claims := token.Claims{
		Id:          auth.PublicID,
		Role:        auth.Role,
		ExpiredTime: time.Duration(cfg.App.ExpiredTime * int(time.Second)),
	}

	token, err := auth.GenerateToken(claims, cfg.App.SecretKey)
	if err != nil {
		slog.ErrorContext(ctx, "[login] error when create token", slog.Any("error", err.Error()))
		return "", "", nil
	}

	return token, auth.Role, nil
}
