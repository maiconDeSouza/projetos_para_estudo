package main

import (
	"api-sistema/internal/config"
	"api-sistema/internal/database"
	"api-sistema/internal/handlers"
	"api-sistema/internal/routes"
	"api-sistema/internal/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}

	services := services.NewServices(db)
	handlers := handlers.NewHandlers(services)
	mux := routes.InitRoutes(handlers)

	log.Printf("Servidor rodando na porta :%s", cfg.SERVER_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", cfg.SERVER_PORT), mux))
}
