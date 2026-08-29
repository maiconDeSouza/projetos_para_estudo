package main

import (
	"api-em-tres-cores/internal/db"
	"api-em-tres-cores/internal/handler"
	"api-em-tres-cores/internal/repository"
	"api-em-tres-cores/internal/router"
	"fmt"
	"log"
	"net/http"
)

const PORT = 2005

func main() {
	// repo := repository.NewInMemoryRepository()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	repoPlayer := repository.NewPostgresPlayerRepository(db)
	repoMatch := repository.NewPostgresMatchRepository(db)

	playerHandler := handler.NewPlayerHandler(repoPlayer)
	matchHandler := handler.NewMatchHandler(repoMatch)

	mux := router.SetupRoutes(playerHandler, matchHandler)

	fmt.Printf("Servidor rodando na porta: %d ... Bora Tricolor ⚽", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
