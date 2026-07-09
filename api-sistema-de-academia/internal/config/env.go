package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_SSLMODE  string
	SERVER_PORT string
}

func LoadEnv() (*ConfigEnv, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := ConfigEnv{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		SERVER_PORT: os.Getenv("SERVER_PORT"),
	}

	return &cfg, nil
}
