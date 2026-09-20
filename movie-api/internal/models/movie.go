package models

import (
	"time"
	"uuid"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:varchar(560);not null" json:"description"`
	Year        string    `gorm:"type:varchar(25);not null" json:"year"`
	Genre       string    `gorm:"type:varchar(60);not null" json:"genre"`
	ImdbID      string    `gorm:"type:varchar(255);not null;uniqueIndex" json:"imdbID"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (m Movie) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&m)
}

type MovieRequestOMDB struct {
	Title    string
	Plot     string
	Year     string
	Genre    string
	ImdbID   string
	Response string
	Error    string
}

type ImdbIDRequest struct {
	ImdbID string `json:"imdbID"`
}

type SearchOMDB struct {
	Title  string
	Year   string
	ImdbID string `json:"imdbID"`
	Poster string
}

type ResultSearchOMDB struct {
	Search       []SearchOMDB
	TotalResults string `json:"totalResults"`
	Response     string
	Error        string
}
