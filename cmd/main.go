package main

import (
	"github.com/taufiqkba/food_app_v2/internal/config"
	"github.com/taufiqkba/food_app_v2/internal/infra/database"
)

func main() {

	err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	cfg := config.GetConfig()

	db, err := database.ConnectPostgres(cfg.DB)
	if err != nil {
		panic(err)
	}
	_ = db
	// fmt.Printf("%+v\n", cfg)
}
