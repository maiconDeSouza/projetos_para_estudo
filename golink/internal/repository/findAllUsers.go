package repository

import (
	"context"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) FindAllUsers() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*domain.User

	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errapp.ErrUserNotFound
	}

	return users, nil
}
