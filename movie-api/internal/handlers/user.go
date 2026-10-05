package handlers

import (
	"encoding/json"
	"movies-api/internal/models"
	"movies-api/internal/services"
	"net/http"
	"net/mail"
)

type HandlersUserInterface interface {
	NewUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type HandlersUser struct {
	services services.ServicesUserInterface
}

func NewHandlersUser(services services.ServicesUserInterface) *HandlersUser {
	return &HandlersUser{services: services}
}

func (h *HandlersUser) NewUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.NewUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		http.Error(w, "Email inválido", http.StatusBadRequest)
		return
	}

	if newUser.Password != newUser.RepeatPassword {
		http.Error(w, "As senhas precisam ser iguais", http.StatusBadRequest)
		return
	}

	user, err := h.services.CreateUser(newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *HandlersUser) Login(w http.ResponseWriter, r *http.Request) {
	login := models.Login{}

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.services.Login(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}
