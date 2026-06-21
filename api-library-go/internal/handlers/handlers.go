package handlers

import (
	"api-library-go/internal/models"
	"api-library-go/internal/services"
	"encoding/json"
	"net/http"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := services.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	services.CreateBook(&newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newBook)
}

func UpBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var upBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&upBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := services.UpBook(&upBook, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
