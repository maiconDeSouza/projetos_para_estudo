package main

import (
	"api-ecommerce/internal/handlers"
	"api-ecommerce/internal/repositories"
	"api-ecommerce/internal/routes"
	"api-ecommerce/internal/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := repositories.NewRepositorie()
	services := services.NewServices(repo)
	handlers := handlers.NewHandlers(services)

	mux := routes.InitRoutes(handlers)

	fmt.Println("Servidor rodando...")
	log.Fatal(http.ListenAndServe(":2005", mux))
}
