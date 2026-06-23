package storage

import (
	"api-library-go-v2/internal/models"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

var mu sync.RWMutex

type Storage struct {
	FilePath string
	IdBook   int
	IdUser   int
	Books    map[int]*models.BookResponse
	Users    map[int]*models.UserResponse
}

type StorageJson struct {
	IdBook int                    `json:"idBook"`
	IdUser int                    `json:"idUser"`
	Books  []*models.BookResponse `json:"books"`
	Users  []*models.UserResponse `json:"users"`
}

func (s *Storage) WriteJSON() {
	sj := StorageJson{
		IdBook: s.IdBook,
		IdUser: s.IdUser,
	}

	for _, v := range s.Books {
		sj.Books = append(sj.Books, v)
	}

	for _, v := range s.Users {
		sj.Users = append(sj.Users, v)
	}

	jsonSJ, err := json.MarshalIndent(sj, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(s.FilePath, jsonSJ, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Storage) ReadJSON() {
	sj := StorageJson{}

	jsonSJ, err := os.ReadFile(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			s.WriteJSON()
			return
		}

		log.Fatal(err)
	}

	err = json.Unmarshal(jsonSJ, &sj)
	if err != nil {
		log.Fatal(err)
	}

	s.IdBook = sj.IdBook
	s.IdUser = sj.IdUser

	for _, book := range sj.Books {
		s.Books[book.ID] = book
	}

	for _, user := range sj.Users {
		s.Users[user.ID] = user
	}
}

func NewStorage() *Storage {
	storage := &Storage{
		FilePath: "./db.json",
		Books:    make(map[int]*models.BookResponse),
		Users:    make(map[int]*models.UserResponse),
	}

	storage.ReadJSON()

	return storage
}

func (s *Storage) CreateNewBook(book *models.BookRequest) *models.BookResponse {
	mu.Lock()
	defer mu.Unlock()

	s.IdBook++

	newBook := &models.BookResponse{
		ID:       s.IdBook,
		Name:     book.Name,
		Author:   book.Author,
		Borrowed: false,
	}

	s.Books[newBook.ID] = newBook

	s.WriteJSON()

	return newBook
}

func (s *Storage) GetAllBooks() map[int]*models.BookResponse {
	mu.RLock()
	defer mu.RUnlock()

	books := make(map[int]*models.BookResponse)

	for k, v := range s.Books {
		books[k] = v
	}

	return books
}

func (s *Storage) GetBook(id int) (*models.BookResponse, error) {
	mu.RLock()
	defer mu.RUnlock()

	book, ok := s.Books[id]
	if !ok {
		return nil, errors.New("Livro não encontrado!")
	}

	return book, nil
}

func (s *Storage) DelBook(id int) (*models.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	book, ok := s.Books[id]
	if !ok {
		return nil, errors.New("Livro não encontrado!")
	}

	delete(s.Books, book.ID)

	s.WriteJSON()

	return book, nil
}
