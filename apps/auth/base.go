package auth

import (
	"database/sql"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/taufiqkba/food_app_v2/apps/auth/login"
)

func InitModule(router *chi.Mux, db *sql.DB) {
	slog.Debug("starting module auth")
	login.Run(router, db)
}
