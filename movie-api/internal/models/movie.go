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
	Year        uint      `gorm:"type:smallint;not null" json:"year"`
	Genre       string    `gorm:"type:varchar(60);not null" json:"genre"`
	ImdbID      string    `gorm:"type:varchar(255);not null" json:"imdbID"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (m Movie) AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&m)
}
