package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	filename := "../../config.test.yaml"
	err := LoadConfig(filename)

	require.Nil(t, err)
	require.NotNil(t, "localhost", cfg.DB.Host)
	require.NotNil(t, "6543", cfg.DB.Port)
	require.NotNil(t, "username", cfg.DB.Username)
	require.NotNil(t, "root", cfg.DB.Password)
	require.NotNil(t, "nb_food_app", cfg.DB.Database)
}
