package models

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CheapSharkID   string         `gorm:"type:varchar(100);not null" json:"cheapshark_id"`
	Title          string         `gorm:"type:varchar(255);not null" json:"title"`
	NormalPrice    float64        `gorm:"type:decimal(10,2);not null" json:"normal_price"`
	SalePrice      float64        `gorm:"type:decimal(10,2);not null" json:"sale_price"`
	SavingsPercent float64        `gorm:"type:decimal(5,2);not null" json:"savings_percent"`
	DealRating     string         `gorm:"type:varchar(50);not null" json:"deal_rating"`
	Thumb          string         `gorm:"type:text" json:"thumb"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete nativo do GORM 🗑️
}
