package repository

import (
	"golink/internal/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	Create(user *domain.User) error
}

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}
