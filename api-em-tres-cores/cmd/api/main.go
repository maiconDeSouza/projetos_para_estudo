package main

import (
	"api-em-tres-cores/internal/handler"
	"api-em-tres-cores/internal/repository"
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome a API em três cores 🇾🇪")
}

const PORT = 2005

func main() {
	repo := repository.NewInMemoryRepository()
	playerHandler := handler.NewPlayerHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler)
	mux.HandleFunc("GET /players", playerHandler.GetPlayers)

	fmt.Printf("Servidor rodando na porta: %d ... Bora Tricolor ⚽", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux))
}
