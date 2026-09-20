package config

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌​ Erro ao carregar o arquivo .env:", err)
	}
}

func InitMUX() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
