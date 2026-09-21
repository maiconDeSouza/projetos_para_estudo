package main

import (
	"fmt"
	"log"
	"movies-api/internal/config"
	"movies-api/internal/database"
	"movies-api/internal/handlers"
	"movies-api/internal/repositories"
	"movies-api/internal/router"
	"movies-api/internal/services"
	"net/http"
	"os"
)

func main() {
	config.InitDotenv()
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	fmt.Println(os.Getenv("OMD_APIKEY"))
	db := database.ConnectDB(dsn)

	mux := config.InitMUX()
	repo := repositories.NewRepo(db)
	services := services.NewServices(repo)
	handlers := handlers.NewHandlers(services)
	router := router.NewRoutes(mux, handlers, os.Getenv("VERSION_API"))
	router.InitRoutes()
	port := os.Getenv("SERVER_PORT")

	fmt.Println("🚀 Servidor iniciado com sucesso na porta " + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
