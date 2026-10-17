package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type contextKey string

const (
	UserSubKey   contextKey = "user_sub"
	UserAdminKey contextKey = "user_admin"
)

type ENV struct {
	SERVER_PORT         string
	MONGO_DB_NAME       string
	MONGO_COLLECTION    string
	MONGO_ROOT_USER     string
	MONGO_ROOT_PASSWORD string
	MONGO_PORT          string
	JWT_SECRET_KEY      string
}

func NewENV() *ENV {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌​ Erro ao carregar o arquivo .env:", err)
	}
	return &ENV{
		SERVER_PORT:         os.Getenv("SERVER_PORT"),
		MONGO_DB_NAME:       os.Getenv("MONGO_DB_NAME"),
		MONGO_COLLECTION:    os.Getenv("MONGO_COLLECTION"),
		MONGO_ROOT_USER:     os.Getenv("MONGO_ROOT_USER"),
		MONGO_ROOT_PASSWORD: os.Getenv("MONGO_ROOT_PASSWORD"),
		MONGO_PORT:          os.Getenv("MONGO_PORT"),
		JWT_SECRET_KEY:      os.Getenv("JWT_SECRET_KEY"),
	}
}
