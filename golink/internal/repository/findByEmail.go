package repository

import (
	"context"
	"errors"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) FindByEmail(email string, user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errapp.ErrUserNotFound
		}
		return err
	}
	return nil
}
