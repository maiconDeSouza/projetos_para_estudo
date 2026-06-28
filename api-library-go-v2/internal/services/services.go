package services

import (
	"api-library-go-v2/internal/models"
	"api-library-go-v2/internal/storage"
	"errors"
	"strconv"
)

type BookUserServices interface {
	CreateBook(book *models.BookRequest) *models.BookResponse
	GetAllBooks() map[int]*models.BookResponse
	GetBook(idString string) (*models.BookResponse, error)
	UpBook(idString string, upBook *models.BookRequest) (*models.BookResponse, error)
	DelBook(idString string) (*models.BookResponse, error)
	CreateUser(user *models.UserRequest) *models.UserResponse
	GetAllUsers() map[int]*models.UserResponse
	GetUser(idString string) (*models.UserResponse, error)
	UpUser(idString string, upUser *models.UserRequest) (*models.UserResponse, error)
	DelUser(idString string) (*models.UserResponse, error)
	BorrowedBook(idUserString, idBookString string) (*models.BorrowedResponse, error)
	ReturnBook(idBookString string) (*models.BookResponse, error)
}

type Services struct {
	storage storage.BookUserRepository
}

func NewServices(storage storage.BookUserRepository) *Services {
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

func (s *Services) CreateUser(user *models.UserRequest) *models.UserResponse {
	return s.storage.CreateUser(user)
}

func (s *Services) GetAllUsers() map[int]*models.UserResponse {
	return s.storage.GetAllUsers()
}

func (s *Services) GetUser(idString string) (*models.UserResponse, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}

	return s.storage.GetUser(id)
}

func (s *Services) UpUser(idString string, upUser *models.UserRequest) (*models.UserResponse, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}

	return s.storage.UpUser(id, upUser)
}

func (s *Services) DelUser(idString string) (*models.UserResponse, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}

	return s.storage.DelUser(id)
}

func (s *Services) BorrowedBook(idUserString, idBookString string) (*models.BorrowedResponse, error) {
	idUser, err := strconv.Atoi(idUserString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}

	idBook, err := strconv.Atoi(idBookString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}

	return s.storage.BorrowedBook(idUser, idBook)
}

func (s *Services) ReturnBook(idBookString string) (*models.BookResponse, error) {
	idBook, err := strconv.Atoi(idBookString)
	if err != nil {
		return nil, errors.New("Parametro tem que ser númerico")
	}

	return s.storage.ReturnBook(idBook)
}
