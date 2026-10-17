package repository

import (
	"context"
	"golink/internal/domain"
	"golink/internal/infra/errapp"
	"time"
	"uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LinkWithOwner junta os dados do Link com o ID do utilizador que o criou
type LinkWithOwner struct {
	UserID uuid.UUID   `bson:"user_id"`
	Data   domain.Link `bson:"data"` // 🟢 CORREÇÃO: Mudado de "links" para "data"
}

// GetLinkByTitle procura um link pelo título e retorna os dados do link e o ID do seu dono
func (r *Repository) GetLinkByTitle(title string) (*LinkWithOwner, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		bson.D{{Key: "unwind", Value: "links"}}, // 🟢 CORREÇÃO: Adicionado o  em unwind

		bson.D{{Key: "$match", Value: bson.M{"links.title": title}}},

		bson.D{{Key: "$project", Value: bson.M{
			"user_id":         "$_id",
			"data._id":        "$links._id",
			"data.title":      "$links.title",
			"data.link":       "$links.link",
			"data.created_at": "$links.created_at",
			"data.update_at":  "$links.update_at",
			"data.count":      "$links.count",
		}}},
	}

	cursor, err := r.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var result LinkWithOwner
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		return &result, nil
	}

	return nil, errapp.ErrNoDocuments
}
