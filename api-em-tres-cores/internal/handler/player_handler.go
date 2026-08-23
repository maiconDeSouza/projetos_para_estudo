package handler

import (
	"api-em-tres-cores/internal/repository"
	"encoding/json"
	"net/http"
)

type PlayerHandler struct {
	repo repository.PlayerRepository
}

func NewPlayerHandler(repo repository.PlayerRepository) *PlayerHandler {
	return &PlayerHandler{repo: repo}
}

func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Erro ao buscar jogador", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(players)
}
