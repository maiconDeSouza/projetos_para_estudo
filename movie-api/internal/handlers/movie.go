package handlers

import (
	"encoding/json"
	"fmt"
	"movies-api/internal/models"
	"movies-api/internal/services"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type HandlersInterface interface {
	NewMovie(w http.ResponseWriter, r *http.Request)
	SearchOMDB(w http.ResponseWriter, r *http.Request)
	AllMovies(w http.ResponseWriter, r *http.Request)
	Movie(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	services services.ServicesInterface
	v        *validator.Validate
}

func NewHandlers(services services.ServicesInterface) *Handlers {
	return &Handlers{services: services, v: validator.New()}
}

func (h *Handlers) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.services.AllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func (h *Handlers) Movie(w http.ResponseWriter, r *http.Request) {
	imdbID := r.PathValue("imdbID")

	movie, err := h.services.Movie(imdbID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (h *Handlers) NewMovie(w http.ResponseWriter, r *http.Request) {
	id := models.ImdbIDRequest{}

	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.v.Struct(id); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMovie, err := h.services.NewMovie(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func (h *Handlers) SearchOMDB(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	fmt.Println(q)
	list, err := h.services.SearchOMDB(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}
