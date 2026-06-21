package storage

import (
	"api-library-go/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

var L = Library{}

func init() {
	L.books = make(map[string]*models.Book)
	L.users = make(map[string]*models.User)

	L.ReadJSON()
}

var pathBooks = "./internal/storage/books.json"
var pathUsers = "./internal/storage/users.json"
var pathID = "./internal/storage/id.json"

type Library struct {
	books  map[string]*models.Book
	users  map[string]*models.User
	idBook int
	idUser int
	gate   sync.RWMutex
}

func (library *Library) WriteJSON() error {
	library.gate.Lock()
	defer library.gate.Unlock()
	var books []*models.Book
	var users []*models.User

	var id = map[string]int{"book": library.idBook, "user": library.idUser}

	for _, v := range library.books {
		books = append(books, v)
	}

	for _, v := range library.users {
		users = append(users, v)
	}

	booksJSON, err := json.MarshalIndent(books, "", " ")
	if err != nil {
		return err
	}

	usersJSON, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	idJSON, err := json.MarshalIndent(id, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(pathBooks, booksJSON, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(pathUsers, usersJSON, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(pathID, idJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (library *Library) ReadJSON() error {
	var books []*models.Book
	var users []*models.User
	var ids map[string]int

	booksJSON, err := os.ReadFile(pathBooks)
	if err != nil {
		return err
	}

	usersJSON, err := os.ReadFile(pathUsers)
	if err != nil {
		return err
	}

	idJSON, err := os.ReadFile(pathID)
	if err != nil {
		return err
	}

	err = json.Unmarshal(booksJSON, &books)
	if err != nil {
		return err
	}

	err = json.Unmarshal(usersJSON, &users)
	if err != nil {
		return err
	}

	err = json.Unmarshal(idJSON, &ids)
	if err != nil {
		return err
	}

	for _, book := range books {
		library.books[book.ID] = book
	}

	for _, user := range users {
		library.users[user.ID] = user
	}

	library.idBook = ids["book"]
	library.idUser = ids["user"]

	return nil
}

func (library *Library) GetAllBooks() map[string]*models.Book {
	return library.books
}

func (library *Library) GetBoook(id string) (*models.Book, error) {
	book, ok := library.books[id]
	if !ok {
		return nil, errors.New("Livro não encontrado")
	}
	return book, nil
}

func (library *Library) AddBook(newBook *models.Book) {
	library.idBook += 1
	newBook.ID = fmt.Sprintf("%d", library.idBook)
	newBook.History = make(map[string]*models.User)
	library.books[newBook.ID] = newBook

	library.WriteJSON()
}

func (library *Library) UpBook(upBook *models.Book, id string) (*models.Book, error) {
	book, exist := library.books[id]
	if !exist {
		return nil, errors.New("Livro não encontrado")
	}

	if upBook.Name != "" {
		book.Name = upBook.Name
	}

	if upBook.Author != "" {
		book.Author = upBook.Author
	}

	library.WriteJSON()
	return book, nil
}

func (library *Library) DelBook(id string) (*models.Book, error) {
	book, exist := library.books[id]
	if !exist {
		return nil, errors.New("Livro não encontrado")
	}

	delete(library.books, id)

	library.WriteJSON()
	return book, nil
}

//Users

func (library *Library) AddUser(newUser *models.User) {
	library.idUser += 1
	newUser.ID = fmt.Sprintf("%d", library.idUser)
	newUser.Book = make(map[string]*models.Book)
	library.users[newUser.ID] = newUser

	library.WriteJSON()
}

func (library *Library) GetAllUsers() map[string]*models.User {
	return library.users
}

func (library *Library) GetUser(id string) (*models.User, error) {
	user, ok := library.users[id]
	if !ok {
		return nil, errors.New("Usuário não encontrado")
	}
	return user, nil
}

func (library *Library) UpUser(upUser *models.User, id string) (*models.User, error) {
	user, exist := library.users[id]
	if !exist {
		return nil, errors.New("Usuário não encontrado")
	}

	if upUser.Name != "" {
		user.Name = upUser.Name
	}

	library.WriteJSON()
	return user, nil
}

func (library *Library) DelUser(id string) (*models.User, error) {
	user, exist := library.users[id]
	if !exist {
		return nil, errors.New("Usuárioo não encontrado")
	}

	delete(library.users, id)

	library.WriteJSON()
	return user, nil
}
