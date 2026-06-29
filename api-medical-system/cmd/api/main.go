package main

import (
	"api-medical-system/internal/handlers"
	"api-medical-system/internal/repositories"
	"api-medical-system/internal/services"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	repositories := repositories.NewRepositories()
	services := services.NewServices(repositories)
	handlers := handlers.NewHandlers(services)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Medical"))
	})

	mux.HandleFunc("POST /api/v1/medical-system/medical", handlers.CreateMedical)

	fmt.Println("Servidor rodando...")
	http.ListenAndServe(":2005", mux)
}
