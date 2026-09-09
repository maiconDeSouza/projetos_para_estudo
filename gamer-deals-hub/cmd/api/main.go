package main

import (
	"log"

	"gamer-deals-hub/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	}

	_, err = database.Connect()
	if err != nil {
		log.Fatalf("Falha crítica ao conectar no banco: %v", err)
	}

	log.Println("Aplicação inicializada com sucesso! 🎮")
}
