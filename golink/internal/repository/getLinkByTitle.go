package repository

import (
	"context"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) GetLinkByTitle(title string) (*domain.Link, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{{Key: "$unwind", Value: "$links"}},

		bson.D{{Key: "$match", Value: bson.M{"links.title": title}}},

		bson.D{{Key: "$project", Value: bson.M{
			"_id":        "$links._id",
			"title":      "$links.title",
			"link":       "$links.link",
			"created_at": "$links.created_at",
			"update_at":  "$links.update_at",
			"count":      "$links.count",
		}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var link domain.Link
		if err := cursor.Decode(&link); err != nil {
			return nil, err
		}
		return &link, nil
	}

	return nil, errapp.ErrNoDocuments
}
