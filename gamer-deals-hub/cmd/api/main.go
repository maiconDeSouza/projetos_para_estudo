package main

import (
	"fmt"
	"gamer-deals-hub/internal/router"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	mux := router.StartServer()
	port := os.Getenv("API_PORT")
	apiPort := fmt.Sprintf(":%s", port)

	fmt.Printf("Servidor rodando na porta :%s\n", port)
	log.Fatal(http.ListenAndServe(apiPort, mux))
}
