package handler

import (
	"api-em-tres-cores/internal/model"
	"api-em-tres-cores/internal/repository"
	"encoding/json"
	"net/http"
)

type MatchHeandler struct {
	repo repository.MatchRepository
}

func NewMatchHandler(repo repository.MatchRepository) *MatchHeandler {
	return &MatchHeandler{repo: repo}
}

func (h *MatchHeandler) GetMatches(w http.ResponseWriter, r *http.Request) {
	matches, err := h.repo.GetAllMatches()
	if err != nil {
		http.Error(w, "Erro ao buscar partidas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matches)
}

func (h *MatchHeandler) CreateMatch(w http.ResponseWriter, r *http.Request) {
	var match model.Match

	err := json.NewDecoder(r.Body).Decode(&match)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.repo.SaveMatch(&match)
	if err != nil {
		http.Error(w, "Erro ao salvar partida", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}
