package repository

import (
	"context"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"
	"uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *Repository) PushLink(userID uuid.UUID, link domain.Link) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"_id":         userID,
		"links.title": bson.M{"$ne": link.Title},
	}

	update := bson.M{
		"$push": bson.M{"links": link},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errapp.ErrUserNotFound
	}

	return nil
}
