package handler

import (
	"api-em-tres-cores/internal/apperr"
	"api-em-tres-cores/internal/model"
	"api-em-tres-cores/internal/repository"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
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
	var req model.MatchRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("02-01-06", req.Date)
	if err != nil {
		http.Error(w, "Formato de data inválido. Use DD-MM-YY (Ex: 24-08-26)", http.StatusBadRequest)
		return
	}

	match := model.Match{
		Opponent:      req.Opponent,
		Date:          parsedDate,
		GoalsSPFC:     req.GoalsSPFC,
		GoalsOpponent: req.GoalsOpponent,
		Events:        []model.Event{},
	}

	err = h.repo.SaveMatch(&match)
	if err != nil {
		var errDuplicate *apperr.ErrDuplicateDate
		if errors.As(err, &errDuplicate) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		http.Error(w, "Erro ao salvar partida", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(match)
}

func (h *MatchHeandler) ResultMatch(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "ID inválido: não é um UUID válido", http.StatusBadRequest)
		return
	}

	var result model.UpdateResult
	err = json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = h.repo.UpdateResult(id, result)
	if errors.Is(err, apperr.ErrNonExistentMatch) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "Erro no servidor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func (h *MatchHeandler) NewEvent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "ID inválido: não é um UUID válido", http.StatusBadRequest)
		return
	}

	var newEvent model.NewEvent
	err = json.NewDecoder(r.Body).Decode(&newEvent)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	event := model.Event{}
	switch {
	case newEvent.EventType == "GOL" && newEvent.GoalsSPFC:
		event.EventType = model.EventGoal
		event.GoalsSPFC = true
	case newEvent.EventType == "GOL":
		event.EventType = model.EventGoal
		event.GoalsSPFC = false
	case newEvent.EventType == "CARTAO_AMARELO":
		event.EventType = model.EventYellowCard
	case newEvent.EventType == "CARTAO_VERMELHO":
		event.EventType = model.EventRedCard
	case newEvent.EventType == "SUBSTITUICAO":
		event.EventType = model.EventReplacement
	default:
		http.Error(w, "Evento inválido", http.StatusBadRequest)
		return

	}

	event.PlayerID = newEvent.PlayerID
	event.Minute = newEvent.Minute
	event.Description = newEvent.Description

	err = h.repo.MatchEvent(id, event)
	if errors.Is(err, apperr.ErrNonExistentMatch) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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
	json.NewEncoder(w).Encode(event)
}
