package services

import (
	"api-library-go-v2/internal/models"
	"api-library-go-v2/internal/storage"
	"errors"
	"strconv"
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

func (s *Services) GetBook(idString string) (*models.BookResponse, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}
	return s.storage.GetBook(id)
}

func (s *Services) UpBook(idString string, upBook *models.BookRequest) (*models.BookResponse, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}
	return s.storage.UpBook(id, upBook)
}

func (s *Services) DelBook(idString string) (*models.BookResponse, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}
	return s.storage.DelBook(id)
}
