package main

import (
	"api-em-tres-cores/internal/handler"
	"api-em-tres-cores/internal/repository"
	"api-em-tres-cores/internal/router"
	"fmt"
	"log"
	"net/http"
)

const PORT = 2005

func main() {
	repo := repository.NewInMemoryRepository()
	playerHandler := handler.NewPlayerHandler(repo)
	mux := router.SetupRoutes(playerHandler)

	fmt.Printf("Servidor rodando na porta: %d ... Bora Tricolor ⚽", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
