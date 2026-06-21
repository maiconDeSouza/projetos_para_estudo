package main

import (
	"api-library-go/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/library/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("oi"))
	})

	mux.HandleFunc("POST /api/v1/library/books", handlers.CreateBook)
	mux.HandleFunc("GET /api/v1/library/books", handlers.GetAllBooks)
	mux.HandleFunc("PUT /api/v1/library/books/{id}", handlers.UpBook)

	fmt.Println("Servidor rodando...")
	http.ListenAndServe(":2005", mux)
}
