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
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	MatchID     uuid.UUID `gorm:"type:uuid;not null;index" json:"match_id"`
	Minute      uint      `gorm:"not null" json:"minute"`
	EventType   EventType `gorm:"type:varchar(20);not null" json:"event_type"`
	PlayerID    uuid.UUID `gorm:"type:uuid;not null;index" json:"player_id"`
	GoalsSPFC   bool      `gorm:"default:true" json:"goals_spfc"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
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
	ID            uuid.UUID `gorm:"type:uuid;primaryKey" json:"id,omitempty"`
	Opponent      string    `gorm:"not null" json:"opponent"`
	Date          time.Time `gorm:"type:date;uniqueIndex;not null" json:"date"`
	GoalsSPFC     uint      `gorm:"default:0" json:"goals_spfc"`
	GoalsOpponent uint      `gorm:"default:0" json:"goals_opponent"`
	Events        []Event   `gorm:"foreignKey:MatchID;constraint:OnDelete:CASCADE" json:"events"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
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
