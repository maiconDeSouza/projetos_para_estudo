package model

import (
	"time"

	"github.com/google/uuid"
)

type Player struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Position  string    `gorm:"not null" json:"position"`
	Number    uint      `gorm:"not null" json:"number"`
	Stats     Stats     `gorm:"embedded" json:"stats"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Stats struct {
	Games         uint `gorm:"default:0" json:"games"`
	MinutesPlayed uint `gorm:"default:0" json:"minutes_played"`
	Goals         uint `gorm:"default:0" json:"goals"`
	Assistis      uint `gorm:"default:0" json:"asssistis"`
	YellowCards   uint `gorm:"default:0" json:"yellow_cards"`
	RedCards      uint `gorm:"default:0" json:"red_cards"`
}
