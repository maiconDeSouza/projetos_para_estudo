package main

import (
	"api-medical-system/internal/handlers"
	"api-medical-system/internal/repositories"
	"api-medical-system/internal/routes"
	"api-medical-system/internal/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repositories := repositories.NewRepositories()
	services := services.NewServices(repositories)
	handlers := handlers.NewHandlers(services)

	mux := routes.InitRoutes(handlers)

	fmt.Println("Servidor rodando...")
	log.Fatal(http.ListenAndServe(":2005", mux))
}
