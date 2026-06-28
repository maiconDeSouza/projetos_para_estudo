package storage

import (
	"api-library-go-v2/internal/models"
	"encoding/json"
	"errors"
	"fmt"
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

type BookUserRepository interface {
	CreateNewBook(book *models.BookRequest) *models.BookResponse
	GetAllBooks() map[int]*models.BookResponse
	GetBook(id int) (*models.BookResponse, error)
	UpBook(id int, upBook *models.BookRequest) (*models.BookResponse, error)
	DelBook(id int) (*models.BookResponse, error)
	CreateUser(user *models.UserRequest) *models.UserResponse
	GetAllUsers() map[int]*models.UserResponse
	GetUser(id int) (*models.UserResponse, error)
	UpUser(id int, upUser *models.UserRequest) (*models.UserResponse, error)
	DelUser(id int) (*models.UserResponse, error)
	BorrowedBook(idUser, idBook int) (*models.BorrowedResponse, error)
	ReturnBook(idBook int) (*models.BookResponse, error)
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

func (s *Storage) UpBook(id int, upBook *models.BookRequest) (*models.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()

	book, ok := s.Books[id]
	if !ok {
		return nil, errors.New("Livro não encontrado!")
	}

	if len(upBook.Name) > 3 {
		book.Name = upBook.Name
	}

	if len(upBook.Author) > 3 {
		book.Author = upBook.Author
	}

	s.WriteJSON()

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

func (s *Storage) CreateUser(user *models.UserRequest) *models.UserResponse {
	mu.Lock()
	defer mu.Unlock()

	s.IdUser++

	newUser := &models.UserResponse{
		ID:    s.IdUser,
		Name:  user.Name,
		Books: make([]string, 0),
	}

	s.Users[newUser.ID] = newUser

	s.WriteJSON()

	return newUser
}

func (s *Storage) GetAllUsers() map[int]*models.UserResponse {
	mu.RLock()
	defer mu.RUnlock()

	users := make(map[int]*models.UserResponse)

	for k, v := range s.Users {
		users[k] = v
	}

	return users
}

func (s *Storage) GetUser(id int) (*models.UserResponse, error) {
	mu.RLock()
	defer mu.RUnlock()
	user, ok := s.Users[id]
	if !ok {
		return nil, errors.New("Usuário não encontrado!")
	}

	return user, nil
}

func (s *Storage) UpUser(id int, upUser *models.UserRequest) (*models.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	user, ok := s.Users[id]
	if !ok {
		return nil, errors.New("Usuário não encontrado!")
	}

	if len(upUser.Name) > 3 {
		user.Name = upUser.Name
	}

	s.WriteJSON()

	return user, nil
}

func (s *Storage) DelUser(id int) (*models.UserResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	user, ok := s.Users[id]
	if !ok {
		return nil, errors.New("Usuário não encontrado!")
	}

	delete(s.Users, user.ID)

	s.WriteJSON()

	return user, nil
}

func (s *Storage) BorrowedBook(idUser, idBook int) (*models.BorrowedResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	response := models.BorrowedResponse{}
	user, ok := s.Users[idUser]
	if !ok {
		return nil, errors.New("Usuário não encontrado!")
	}

	book, ok := s.Books[idBook]
	if !ok {
		return nil, errors.New("Livro não encontrado!")
	}

	if book.Borrowed {
		return nil, errors.New("Livro já emprestado")
	}

	response.NameUser = user.Name
	response.NameBook = book.Name

	book.Borrowed = true
	user.Books = append(user.Books, fmt.Sprintf("%s", response.NameBook))

	s.WriteJSON()

	return &response, nil
}

func (s *Storage) ReturnBook(idBook int) (*models.BookResponse, error) {
	mu.Lock()
	defer mu.Unlock()
	book, ok := s.Books[idBook]
	if !ok {
		return nil, errors.New("Livro não encontrado!")
	}

	if !book.Borrowed {
		return nil, errors.New("Livro já está na biblioteca!")
	}

	book.Borrowed = false

	s.WriteJSON()

	return book, nil
}
