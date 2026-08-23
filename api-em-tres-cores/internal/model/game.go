package model

import "time"

type EventType string

const (
	EventGoal        EventType = "GOL"
	EventYellowCard  EventType = "CARTAO_AMARELO"
	EventRedCard     EventType = "CARTAO_VERMELHO"
	EventReplacement EventType = "SUBSTITUICAO"
)

type Event struct {
	ID          [16]byte  `json:"id"`
	Minute      uint      `json:"minute"`
	EventType   EventType `json:"event_type"`
	PlayerID    [16]byte  `json:"player_id"`
	GoalsSPFC   bool      `json:"goals_spfc"`
	Description string    `json:"description"`
}

type Match struct {
	ID            [16]byte  `json:"id"`
	Opponent      string    `json:"opponent"`
	Date          time.Time `json:"date"`
	GoalsSPFC     uint      `json:"goals_spfc"`
	GoalsOpponent uint      `json:"goals_opponent"`
	Events        []Event   `json:"events"`
}
