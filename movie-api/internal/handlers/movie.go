package handlers

import (
	"encoding/json"
	"fmt"
	"movies-api/internal/models"
	"movies-api/internal/services"
	"net/http"
	"uuid"

	"github.com/go-playground/validator/v10"
)

type HandlersMovieInterface interface {
	NewMovie(w http.ResponseWriter, r *http.Request)
	SearchOMDB(w http.ResponseWriter, r *http.Request)
	AllMovies(w http.ResponseWriter, r *http.Request)
	Movie(w http.ResponseWriter, r *http.Request)
}

type HandlersMovie struct {
	services services.ServicesMovieInterface
	v        *validator.Validate
}

func NewHandlersMovie(services services.ServicesMovieInterface) *HandlersMovie {
	return &HandlersMovie{services: services, v: validator.New()}
}

func (h *HandlersMovie) AllMovies(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("user_sub").(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movies, err := h.services.AllMovies(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

func (h *HandlersMovie) Movie(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("user_sub").(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	imdbID := r.PathValue("imdbID")

	movie, err := h.services.Movie(id, imdbID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}

func (h *HandlersMovie) NewMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.Context().Value("user_sub").(string)
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	idMovie := models.ImdbIDRequest{}

	if err := json.NewDecoder(r.Body).Decode(&idMovie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.v.Struct(idMovie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newMovie, err := h.services.NewMovie(id, idMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func (h *HandlersMovie) SearchOMDB(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	fmt.Println(q)
	list, err := h.services.SearchOMDB(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)
}
