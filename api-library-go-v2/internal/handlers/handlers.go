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

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	book, err := h.service.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *Handler) UpBook(w http.ResponseWriter, r *http.Request) {
	var upBook models.BookRequest
	id := r.PathValue("id")

	err := json.NewDecoder(r.Body).Decode(&upBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := h.service.UpBook(id, &upBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *Handler) DelBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	book, err := h.service.DelBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.UserRequest{}

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := h.service.CreateUser(&newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	user, err := h.service.GetUser(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
