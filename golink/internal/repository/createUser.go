package repository

import (
	"context"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errapp.ErrEmailAlreadyInUse
		}
		return err
	}

	return nil
}
