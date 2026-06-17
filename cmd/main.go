package main

import (
	"log/slog"
	"os"

	"github.com/taufiqkba/food_app_v2/internal/config"
	"github.com/taufiqkba/food_app_v2/internal/server"
)

func main() {

	err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	log := slog.New(logHandler)
	slog.SetDefault(log)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
