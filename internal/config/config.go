package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App AppConfig `yaml:"app"`
	DB  DBConfig  `yaml:"db"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSLMode  string `yaml:"sslmode"`

	// for connection pool
	MaxIdleConnection int `yaml:"max_idle_connection"`
	MaxOpenConnection int `yaml:"max_open_connection"`
	MaxLifetime       int `yaml:"max_lifetime"`
	MaxIdleTime       int `yaml:"max_idle_time"`
}

var cfg *Config = &Config{}

func LoadConfig(filename string) error {

	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(fileBytes, cfg)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return cfg
}
