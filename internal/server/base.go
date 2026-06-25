package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/taufiqkba/food_app_v2/apps/auth"
	"github.com/taufiqkba/food_app_v2/apps/employee"
	"github.com/taufiqkba/food_app_v2/internal/config"
	"github.com/taufiqkba/food_app_v2/internal/infra/database"
)

func Start() error {
	cfg := config.GetConfig()

	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		return err
	}

	// router
	router := chi.NewRouter()

	// add router
	auth.InitModule(router, db)
	employee.InitModule(router, db)

	slog.Info("server "+cfg.App.Name+" starting", slog.String("port", cfg.App.Port))
	http.ListenAndServe(":"+cfg.App.Port, router)
	return nil
}
