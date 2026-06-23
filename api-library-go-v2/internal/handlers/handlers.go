package handlers

import (
	"api-library-go-v2/internal/models"
	"api-library-go-v2/internal/services"
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{service: services}
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.BookRequest

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newBook := h.service.CreateBook(&book)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := h.service.GetAllBooks()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
