package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/taufiqkba/food_app_v2/internal/config"
)

func ConnectPostgres(cfg config.DBConfig) (db *sql.DB, err error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&search_path=nb_food_app",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode,
	)
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConnection)
	db.SetMaxIdleConns(cfg.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleTime) * time.Second)

	return
}
