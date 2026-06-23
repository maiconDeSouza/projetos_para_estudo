package storage

import (
	"api-library-go-v2/internal/models"
	"encoding/json"
	"log"
	"os"
	"sync"
)

var mu sync.RWMutex

type Storage struct {
	filePath string
	idBook   int
	idUser   int
	books    map[int]*models.BookResponse
	users    map[int]*models.UserResponse
}

type StorageJson struct {
	idBook int
	idUser int
	books  []*models.BookResponse
	users  []*models.UserResponse
}

func (s *Storage) WriteJSON() {
	sj := StorageJson{}

	sj.idBook = s.idBook
	sj.idUser = s.idUser

	for _, v := range s.books {
		sj.books = append(sj.books, v)
	}

	for _, v := range s.users {
		sj.users = append(sj.users, v)
	}

	jsonSJ, err := json.MarshalIndent(sj, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(s.filePath, jsonSJ, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Storage) ReadJson() {
	sj := StorageJson{}
	jsonSJ, err := os.ReadFile(s.filePath)
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

	s.idBook = sj.idBook
	s.idUser = sj.idUser

	for _, book := range sj.books {
		s.books[book.ID] = book
	}

	for _, user := range sj.users {
		s.users[user.ID] = user
	}
}

func NewStorage() *Storage {
	storage := Storage{
		filePath: "./db.json",
	}

	storage.ReadJson()

	return &storage

}

func (s *Storage) CreateNewBook(book *models.BookRequest) *models.BookResponse {
	id := len(s.books) + 1

	newBook := models.BookResponse{
		ID:       id,
		Name:     book.Name,
		Author:   book.Author,
		Borrowed: false,
	}

	s.idBook = id
	s.books[id] = &newBook

	s.WriteJSON()

	return &newBook
}

func (s *Storage) GetAllBooks() *map[int]*models.BookResponse {
	return &s.books
}
