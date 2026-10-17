package repository

import (
	"golink/internal/domain"
	"uuid"

	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	Create(user *domain.User) error
	FindByID(id uuid.UUID, user *domain.User) error
	FindByEmail(email string, user *domain.User) error
	FindAllUsers() ([]*domain.User, error)
	PushLink(userID uuid.UUID, link domain.Link) error
	GetLinkByTitle(title string) (*LinkWithOwner, error)
	UpdateUser(id uuid.UUID, user *domain.User) error
}

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}
