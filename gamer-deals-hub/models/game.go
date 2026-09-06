package models

import (
	"gorm.io/gorm"
)

// Game representa um jogo cadastrado no sistema
type Game struct {
	gorm.Model
	Title         string  `json:"title"`
	CheapSharkID  string  `json:"cheap_shark_id" gorm:"uniqueIndex"`
	SteamAppID    string  `json:"steam_app_id"`
	CheapestPrice float64 `json:"cheapest_price"`
	Wishlisted    bool    `json:"wishlisted" gorm:"default:true"`
	Deals         []Deal  `json:"deals,omitempty" gorm:"foreignKey:GameID"`
}

// Deal representa uma oferta encontrada para o jogo
type Deal struct {
	gorm.Model
	GameID      uint    `json:"game_id"`
	StoreID     string  `json:"store_id"`
	Price       float64 `json:"price"`
	RetailPrice float64 `json:"retail_price"`
	Savings     float64 `json:"savings"`             // Porcentagem de desconto
	CostBenefit string  `json:"cost_benefit_rating"` // Nota calculada pela nossa API (ex: "Excelente", "Regular")
	DealID      string  `json:"deal_id"`
}
