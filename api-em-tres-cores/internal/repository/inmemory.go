package repository

import (
	"api-em-tres-cores/internal/model"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type InMemoryRepository struct {
	mu      sync.RWMutex
	players map[uuid.UUID]model.Player
	matches map[uuid.UUID]model.Match
}

func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		players: make(map[uuid.UUID]model.Player),
		matches: make(map[uuid.UUID]model.Match),
	}

	repo.players[uuid.New()] = model.Player{
		ID:       uuid.New(),
		Name:     "Calleri",
		Position: "Atacante",
		Number:   9,
		Stats:    model.Stats{Games: 10, MinutesPlayed: 850, Goals: 6, Assistis: 1, YellowCards: 0, RedCards: 0},
	}

	repo.players[uuid.New()] = model.Player{
		ID:       uuid.New(),
		Name:     "Lucas Moura",
		Position: "Meia",
		Number:   7,
		Stats:    model.Stats{Games: 8, MinutesPlayed: 710, Goals: 4, Assistis: 10, YellowCards: 1, RedCards: 0},
	}
	return repo
}

func (r *InMemoryRepository) GetAll() ([]model.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.Player

	for _, player := range r.players {
		list = append(list, player)
	}

	return list, nil
}

func (r *InMemoryRepository) GetByID(id uuid.UUID) (*model.Player, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	player, exists := r.players[id]
	if !exists {
		return nil, fmt.Errorf("jogador não encontrado")
	}

	return &player, nil
}

func (r *InMemoryRepository) Create(player *model.Player) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	player.ID = uuid.New()
	r.players[player.ID] = *player
	return nil
}

func (r *InMemoryRepository) UpdateStats(playerID uuid.UUID, minutes uint, eventType model.EventType) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	player, exists := r.players[playerID]
	if !exists {
		return fmt.Errorf("jogador não encontrado")
	}

	player.Stats.MinutesPlayed += minutes

	switch eventType {
	case model.EventGoal:
		player.Stats.Goals++
	case model.EventYellowCard:
		player.Stats.YellowCards++
	case model.EventRedCard:
		player.Stats.RedCards++
	}

	r.players[playerID] = player
	return nil
}

func (r *InMemoryRepository) SaveMatch(match *model.Match) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	match.ID = uuid.New()
	r.matches[match.ID] = *match
	return nil
}

func (r *InMemoryRepository) GetMatchByID(id uuid.UUID) (*model.Match, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	match, exists := r.matches[id]
	if !exists {
		return nil, fmt.Errorf("partida não encontrada")
	}
	return &match, nil
}

func (r *InMemoryRepository) GetAllMatches() ([]model.Match, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var list []model.Match
	for _, match := range r.matches {
		list = append(list, match)
	}
	return list, nil
}
