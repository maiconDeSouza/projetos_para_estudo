package repository

import (
	"gamer-deals-hub/internal/models"

	"gorm.io/gorm"
)

type WishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{db: db}
}

func (r *WishlistRepository) Save(game *models.Game) error {
	return r.db.Create(game).Error
}

func (r *WishlistRepository) FindAll() ([]models.Game, error) {
	var games []models.Game
	err := r.db.Find(&games).Error
	return games, err
}
