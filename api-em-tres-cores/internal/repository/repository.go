package repository

import (
	"api-em-tres-cores/internal/model"

	"github.com/google/uuid"
)

type PlayerRepository interface {
	GetPlayerByID(id uuid.UUID) (*model.Player, error)
	GetAllPlayers() ([]model.Player, error)
	CreatePlayer(player *model.Player) error
	UpdateStatsPlayer(playerID uuid.UUID, minute uint, eventType model.EventType) error
}

type MatchRepository interface {
	SaveMatch(match *model.Match) error
	GetMatchByID(id uuid.UUID) (*model.Match, error)
	GetAllMatches() ([]model.Match, error)
	UpdateResult(matchID uuid.UUID, result model.UpdateResult) error
	MatchEvent(matchID uuid.UUID, event model.Event) error
}
