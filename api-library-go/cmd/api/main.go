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
	mux.HandleFunc("GET /api/v1/library/books/{id}", handlers.GetBook)
	mux.HandleFunc("PUT /api/v1/library/books/{id}", handlers.UpBook)
	mux.HandleFunc("DELETE /api/v1/library/books/{id}", handlers.DelBook)

	mux.HandleFunc("POST /api/v1/library/users", handlers.CreatUser)
	mux.HandleFunc("GET /api/v1/library/users", handlers.GetAllUsers)
	mux.HandleFunc("GET /api/v1/library/users/{id}", handlers.GetUser)
	mux.HandleFunc("PUT /api/v1/library/users/{id}", handlers.UpUser)
	mux.HandleFunc("DELETE /api/v1/library/users/{id}", handlers.DelUser)

	fmt.Println("Servidor rodando...")
	http.ListenAndServe(":2005", mux)
}
