package repository

import (
	"api-em-tres-cores/internal/model"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	GetByID(id uuid.UUID) (*model.Player, error)
	GetAll() ([]model.Player, error)
	Create(player *model.Player) error
	UpdateStats(playerID uuid.UUID, minute uint, eventType model.EventType) error
}

type MatchRepository interface {
	SaveMatch(match *model.Match) error
	GetMatchByID(id uuid.UUID) (*model.Match, error)
	GetAllMatches() ([]model.Match, error)
}
