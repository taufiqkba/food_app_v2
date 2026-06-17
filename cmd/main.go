package main

import (
	"fmt"

	"github.com/taufiqkba/food_app_v2/internal/config"
)

func main() {

	err := config.LoadConfig("./config.yaml")
	if err != nil {
		panic(err)
	}

	cfg := config.GetConfig()
	fmt.Printf("%+v\n", cfg)
}
