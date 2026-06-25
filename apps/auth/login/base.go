package login

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB) {
	slog.Debug("run service login", slog.Any("path", "/api/v1/auth/login"), slog.Any("method", http.MethodPost))
	repo := NewRepositoryLogin(db)
	service := newServiceLogin(repo)
	handler := newHandlerLogin(service)

	router.Post("/api/v1/auth/login", handler.login)

}
