package handler

import (
	"api-em-tres-cores/internal/apperr"
	"api-em-tres-cores/internal/model"
	"api-em-tres-cores/internal/repository"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type PlayerHandler struct {
	repo repository.PlayerRepository
}

func NewPlayerHandler(repo repository.PlayerRepository) *PlayerHandler {
	return &PlayerHandler{repo: repo}
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.repo.GetAllPlayers()
	if err != nil {
		http.Error(w, "Erro ao buscar jogador", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(players)
}

func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	var player model.Player

	err := json.NewDecoder(r.Body).Decode(&player)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.repo.CreatePlayer(&player)
	if err != nil {
		var errDuplicate *apperr.ErrDuplicatePlayer
		if errors.As(err, &errDuplicate) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, "Erro ao salvar jogador", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(player)
}

func (h *PlayerHandler) PlayingTime(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "ID inválido: não é um UUID válido", http.StatusBadRequest)
		return
	}

	var playingTime model.PlayingTime
	err = json.NewDecoder(r.Body).Decode(&playingTime)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateStatsPlayer(id, playingTime.Minutes, model.EventPlayingTime)
	if errors.Is(err, apperr.ErrNonExistentPlayer) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(playingTime)
}
