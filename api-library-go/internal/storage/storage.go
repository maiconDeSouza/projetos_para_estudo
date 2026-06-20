package storage

import "api-library-go/internal/models"

type Library struct {
	books map[string]*models.Book
	users map[string]*models.User
}
