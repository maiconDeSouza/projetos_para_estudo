package main

import (
	"fmt"
	"log"

	"gamer-deals-hub/internal/database"
	"gamer-deals-hub/internal/repository"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: Arquivo .env não encontrado. Usando variáveis de ambiente do sistema.")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Falha crítica ao conectar no banco: %v", err)
	}

	wishlistRepository := repository.NewWishlistRepository(db)
	fmt.Println(wishlistRepository)

	log.Println("Aplicação inicializada com sucesso! 🎮")
}
