package repository

import (
	"api-em-tres-cores/internal/apperr"
	"api-em-tres-cores/internal/model"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostgresMatchRepository struct {
	db *gorm.DB
}

func NewPostgresMatchRepository(db *gorm.DB) *PostgresMatchRepository {
	return &PostgresMatchRepository{db: db}
}

func (r *PostgresMatchRepository) SaveMatch(match *model.Match) error {
	match.ID = uuid.New()
	result := r.db.Create(match)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresMatchRepository) GetMatchByID(id uuid.UUID) (*model.Match, error) {
	var match model.Match

	result := r.db.Preload("Events").First(&match, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperr.ErrNonExistentMatch
		}
		return nil, result.Error
	}

	return &match, nil
}

func (r *PostgresMatchRepository) GetAllMatches() ([]model.Match, error) {
	var matches []model.Match

	result := r.db.Preload("Events").Find(&matches)
	if result.Error != nil {
		return nil, result.Error
	}

	return matches, nil
}

func (r *PostgresMatchRepository) UpdateResult(matchID uuid.UUID, result model.UpdateResult) error {
	var match model.Match

	query := r.db.First(&match, "id = ?", matchID)
	if query.Error != nil {
		if errors.Is(query.Error, gorm.ErrRecordNotFound) {
			return apperr.ErrNonExistentMatch
		}
		return query.Error
	}

	match.GoalsSPFC = result.GoalsSPFC
	match.GoalsOpponent = result.GoalsOpponent

	if err := r.db.Save(&match).Error; err != nil {
		return fmt.Errorf("erro ao atualizar resultado do jogo: %w", err)
	}

	return nil

}

func (r *PostgresMatchRepository) MatchEvent(matchID uuid.UUID, event model.Event) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var match model.Match
		if err := tx.First(&match, "id = ?", matchID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperr.ErrNonExistentMatch
			}
			return err
		}

		event.ID = uuid.New()
		event.MatchID = matchID
		if err := tx.Create(&event).Error; err != nil {
			return fmt.Errorf("erro ao criar evento: %w", err)
		}

		playerRepo := NewPostgresPlayerRepository(tx)
		if err := playerRepo.UpdateStatsPlayer(event.PlayerID, event.Minute, event.EventType); err != nil {
			return err
		}

		return nil
	})
}
