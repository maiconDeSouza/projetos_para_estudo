package model

import (
	"uuid"
)

type Game struct {
	External          string    `json:"external" gorm:"type:varchar(255); not null;column:external"`
	Cheapest          float64   `json:"cheapest" gorm:"type:decimal(10,2); not null;column:cheapest"`
	CheapestPriceEver float64   `json:"cheapestPriceEver" gorm:"type:decimal(10,2);not null;column:cheapestPriceEver"`
	MostExpensive     float64   `json:"mostExpensive" gorm:"type:decimal(10,2);not null;column:mostExpensive"`
	GameID            string    `json:"gameID" gorm:"primaryKey;column:gameID"`
	Thumb             string    `json:"thumb" gorm:"type:varchar(512);column:thumb"`
	UserID            uuid.UUID `json:"userID" gorm:"type:uuid;column:userID"`
}
