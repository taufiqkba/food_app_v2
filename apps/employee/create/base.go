package create

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(router *chi.Mux, db *sql.DB) {
	slog.Debug("run service create employee", slog.Any("path", "/api/v1/employees"), slog.Any("method", http.MethodPost))

	repo := NewCreateEmployeeRepository(db)
	service := NewCreateEmployeeService(repo)
	handler := NewHandlerCreateEmployee(service)

	router.Post("/api/v1/employees", handler.createEmployee)

}
