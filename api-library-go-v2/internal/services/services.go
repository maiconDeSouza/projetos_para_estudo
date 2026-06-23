package services

import (
	"api-library-go-v2/internal/models"
	"api-library-go-v2/internal/storage"
)

type Services struct {
	storage *storage.Storage
}

func NewServices(storage *storage.Storage) *Services {
	return &Services{
		storage: storage,
	}
}

func (s *Services) CreateBook(book *models.BookRequest) *models.BookResponse {
	return s.storage.CreateNewBook(book)
}

func (s *Services) GetAllBooks() map[int]*models.BookResponse {
	return s.storage.GetAllBooks()
}
