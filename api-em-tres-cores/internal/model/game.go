package model

import (
	"time"

	"github.com/google/uuid"
)

type EventType string

const (
	EventGoal        EventType = "GOL"
	EventYellowCard  EventType = "CARTAO_AMARELO"
	EventRedCard     EventType = "CARTAO_VERMELHO"
	EventReplacement EventType = "SUBSTITUICAO"
	EventPlayingTime EventType = "TEM_EM_CAMPO"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Minute      uint      `json:"minute"`
	EventType   EventType `json:"event_type"`
	PlayerID    uuid.UUID `json:"player_id"`
	GoalsSPFC   bool      `json:"goals_spfc"`
	Description string    `json:"description"`
}

type NewEvent struct {
	PlayerID    uuid.UUID `json:"player_id"`
	Minute      uint      `json:"minute"`
	EventType   EventType `json:"event_type"`
	GoalsSPFC   bool      `json:"goals_spfc"`
	Description string    `json:"description"`
}

type PlayingTime struct {
	Minutes uint `json:"minutes"`
}

type Match struct {
	ID            uuid.UUID `json:"id,omitempty"`
	Opponent      string    `json:"opponent"`
	Date          time.Time `json:"date"`
	GoalsSPFC     uint      `json:"goals_spfc"`
	GoalsOpponent uint      `json:"goals_opponent"`
	Events        []Event   `json:"events"`
}

type UpdateResult struct {
	GoalsSPFC     uint `json:"goals_spfc"`
	GoalsOpponent uint `json:"goals_opponent"`
}

type MatchRequest struct {
	Opponent      string `json:"opponent"`
	Date          string `json:"date"`
	GoalsSPFC     uint   `json:"goals_spfc,omitempty"`
	GoalsOpponent uint   `json:"goals_opponent,omitempty"`
}
