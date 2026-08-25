package repository

import (
	"api-em-tres-cores/internal/apperr"
	"api-em-tres-cores/internal/model"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type InMemoryRepository struct {
	mu      sync.RWMutex
	players map[uuid.UUID]model.Player
	matches map[uuid.UUID]*model.Match
}

func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		players: make(map[uuid.UUID]model.Player),
		matches: make(map[uuid.UUID]*model.Match),
	}

	id := uuid.New()
	repo.players[id] = model.Player{
		ID:       id,
		Name:     "Calleri",
		Position: "Atacante",
		Number:   9,
		Stats:    model.Stats{Games: 10, MinutesPlayed: 850, Goals: 6, Assistis: 1, YellowCards: 0, RedCards: 0},
	}

	id = uuid.New()
	repo.players[id] = model.Player{
		ID:       id,
		Name:     "Lucas Moura",
		Position: "Meia",
		Number:   7,
		Stats:    model.Stats{Games: 8, MinutesPlayed: 710, Goals: 4, Assistis: 10, YellowCards: 1, RedCards: 0},
	}
	return repo
}

func (r *InMemoryRepository) GetAllPlayers() ([]model.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.Player

	for _, player := range r.players {
		list = append(list, player)
	}

	return list, nil
}

func (r *InMemoryRepository) GetPlayerByID(id uuid.UUID) (*model.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	player, exists := r.players[id]
	if !exists {
		return nil, fmt.Errorf("jogador não encontrado")
	}

	return &player, nil
}

func (r *InMemoryRepository) CreatePlayer(newPlayer *model.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, player := range r.players {
		if player.Number == newPlayer.Number {
			return &apperr.ErrDuplicatePlayer{Name: player.Name, Number: player.Number}
		}
	}

	newPlayer.ID = uuid.New()
	r.players[newPlayer.ID] = *newPlayer
	return nil
}

func (r *InMemoryRepository) UpdateStatsPlayer(playerID uuid.UUID, minutes uint, eventType model.EventType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	player, exists := r.players[playerID]
	if !exists {
		return apperr.ErrNonExistentPlayer
	}

	switch eventType {
	case model.EventGoal:
		player.Stats.Goals++
	case model.EventYellowCard:
		player.Stats.YellowCards++
	case model.EventRedCard:
		player.Stats.RedCards++
	case model.EventPlayingTime:
		player.Stats.MinutesPlayed += minutes
	}

	r.players[playerID] = player
	return nil
}

func (r *InMemoryRepository) SaveMatch(newMatch *model.Match) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, match := range r.matches {
		if newMatch.Date.Equal(match.Date) {
			return &apperr.ErrDuplicateDate{Date: newMatch.Date}
		}
	}

	newMatch.ID = uuid.New()
	r.matches[newMatch.ID] = newMatch
	return nil
}

func (r *InMemoryRepository) UpdateResult(matchID uuid.UUID, result model.UpdateResult) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	match, exists := r.matches[matchID]
	if !exists {
		return apperr.ErrNonExistentMatch
	}

	match.GoalsSPFC = result.GoalsSPFC
	match.GoalsOpponent = result.GoalsOpponent

	r.matches[matchID] = match

	return nil
}

func (r *InMemoryRepository) GetMatchByID(id uuid.UUID) (*model.Match, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	match, exists := r.matches[id]
	if !exists {
		return nil, apperr.ErrNonExistentMatch
	}
	return match, nil
}

func (r *InMemoryRepository) GetAllMatches() ([]model.Match, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.Match
	for _, match := range r.matches {
		list = append(list, *match)
	}
	return list, nil
}

func (r *InMemoryRepository) MatchEvent(matchID uuid.UUID, event model.Event) error {
	match, err := r.GetMatchByID(matchID)
	if err != nil {
		return err
	}
	event.ID = uuid.New()

	err = r.UpdateStatsPlayer(event.PlayerID, event.Minute, event.EventType)
	if err != nil {
		return err
	}

	match.Events = append(match.Events, event)

	return nil
}
