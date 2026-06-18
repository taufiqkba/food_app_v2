package login

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	slog.Debug("run service login", slog.Any("path", "/v1/auth/login"))
	repo := NewRepositoryLogin(db)
	service := newServiceLogin(repo)
	handler := newHandlerLogin(service)

	router.Post("/v1/auth/login", handler.login)

}
