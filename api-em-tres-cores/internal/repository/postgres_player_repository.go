package repository

import (
	"api-em-tres-cores/internal/apperr"
	"api-em-tres-cores/internal/model"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresPlayerRepository struct {
	db *gorm.DB
}

func NewPostgresPlayerRepository(db *gorm.DB) *PostgresPlayerRepository {
	return &PostgresPlayerRepository{db: db}
}

func (r *PostgresPlayerRepository) GetAllPlayers() ([]model.Player, error) {
	var players []model.Player

	result := r.db.Find(&players)
	if result.Error != nil {
		return nil, fmt.Errorf("erro ao buscar jogadores: %w", result.Error)
	}

	return players, nil
}

func (r *PostgresPlayerRepository) GetPlayerByID(id uuid.UUID) (*model.Player, error) {
	var player model.Player

	result := r.db.Find(&player, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("jogador não encontrado")
		}
		return nil, result.Error
	}
	return &player, nil
}

func (r *PostgresPlayerRepository) CreatePlayer(newPlayer *model.Player) error {
	var existingPlayer model.Player
	err := r.db.Where("number = ?", newPlayer.Number).First(&existingPlayer).Error
	if err == nil {
		return &apperr.ErrDuplicatePlayer{Name: existingPlayer.Name, Number: existingPlayer.Number}
	}

	newPlayer.ID = uuid.New()
	result := r.db.Create(newPlayer)
	if result.Error != nil {
		return fmt.Errorf("erro ao criar jogador: %w", result.Error)
	}
	return nil
}

func (r *PostgresPlayerRepository) UpdateStatsPlayer(playerID uuid.UUID, minutes uint, eventType model.EventType) error {
	var player model.Player

	result := r.db.First(&player, playerID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return apperr.ErrNonExistentPlayer
		}
		return result.Error
	}

	switch eventType {
	case model.EventGoal:
		player.Stats.Goals++
	case model.EventYellowCard:
		player.Stats.YellowCards++
	case model.EventRedCard:
		player.Stats.RedCards++
	case model.EventPlayingTime:
		player.Stats.MinutesPlayed += minutes
	}

	if err := r.db.Save(&player).Error; err != nil {
		return fmt.Errorf("erro ao atualizar estatísticas do jogador: %w", err)
	}

	return nil
}
