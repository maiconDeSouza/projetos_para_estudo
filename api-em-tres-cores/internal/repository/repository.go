package repository

import "api-em-tres-cores/internal/model"

type PlayerRepository interface {
	GetByID(id [16]byte) (*model.Player, error)
	GetAll() ([]model.Player, error)
	Create(player *model.Player) error
	UpdateStats(playerID [16]byte, minute uint, eventType model.EventType) error
}

type MatchRepository interface {
	SaveMatch(match *model.Match) error
	GetMatchByID(id [16]byte) (*model.Match, error)
	GetAllMatches() ([]model.Match, error)
}
