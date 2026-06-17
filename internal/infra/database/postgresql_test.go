package database

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taufiqkba/food_app_v2/internal/config"
)

func init() {
	err := config.LoadConfig("../../../config.yaml")
	if err != nil {
		panic(err)
	}
}
func TestConnectPostgres(t *testing.T) {
	cfg := config.GetConfig()

	db, err := ConnectPostgres(cfg.DB)
	require.Nil(t, err)
	require.NotNil(t, db)
}
