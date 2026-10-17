package repository

import (
	"context"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"
	"uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) UpdateUser(id uuid.UUID, user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": user.ID}

	result, err := r.collection.ReplaceOne(ctx, filter, user)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errapp.ErrUserNotFound
	}

	return nil
}
