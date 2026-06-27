package main

import (
	"api-library-go-v2/internal/handlers"
	"api-library-go-v2/internal/services"
	"api-library-go-v2/internal/storage"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	storage := storage.NewStorage()
	services := services.NewServices(storage)
	handlers := handlers.NewHandler(services)

	mux.HandleFunc("GET /api/v2/library/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oi"))
	})

	mux.HandleFunc("POST /api/v2/library/books", handlers.CreateBook)
	mux.HandleFunc("GET /api/v2/library/books", handlers.GetAllBooks)
	mux.HandleFunc("GET /api/v2/library/books/{id}", handlers.GetBook)
	mux.HandleFunc("PUT /api/v2/library/books/{id}", handlers.UpBook)
	mux.HandleFunc("DELETE /api/v2/library/books/{id}", handlers.DelBook)

	mux.HandleFunc("POST /api/v2/library/users", handlers.CreateUser)
	mux.HandleFunc("GET /api/v2/library/users", handlers.GetAllUsers)
	mux.HandleFunc("GET /api/v2/library/users/{id}", handlers.GetUser)
	mux.HandleFunc("PUT /api/v2/library/users/{id}", handlers.UpUser)
	mux.HandleFunc("DELETE /api/v2/library/users/{id}", handlers.DelUser)

	mux.HandleFunc("GET /api/v2/library/users/{idUser}/{idBook}", handlers.BorrowedBook)
	mux.HandleFunc("GET /api/v2/library/users/return/{idBook}", handlers.ReturnBook)

	fmt.Println("servidor rodando...")
	http.ListenAndServe(":2005", mux)
}
