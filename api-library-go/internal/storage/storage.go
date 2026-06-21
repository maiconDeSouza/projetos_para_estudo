package storage

import (
	"api-library-go/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var L = Library{}

func init() {
	L.books = make(map[string]*models.Book)
	L.users = make(map[string]*models.User)

	L.ReadJSON()
}

var pathBooks = "./internal/storage/books.json"
var pathUsers = "./internal/storage/users.json"

type Library struct {
	books map[string]*models.Book
	users map[string]*models.User
}

func (library *Library) WriteJSON() error {
	var books []*models.Book
	var users []*models.User

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

	err = os.WriteFile(pathBooks, booksJSON, 0644)
	if err != nil {
		return err
	}

	err = os.WriteFile(pathUsers, usersJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (library *Library) ReadJSON() error {
	var books []*models.Book
	var users []*models.User

	booksJSON, err := os.ReadFile(pathBooks)
	if err != nil {
		return err
	}

	usersJSON, err := os.ReadFile(pathUsers)
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

	for _, book := range books {
		library.books[book.ID] = book
	}

	for _, user := range users {
		library.users[user.ID] = user
	}

	return nil
}

func (library *Library) GetAllBooks() map[string]*models.Book {
	return library.books
}

func (library *Library) AddBook(newBook *models.Book) {
	id := fmt.Sprintf("%d", len(library.books)+1)
	newBook.ID = id
	newBook.History = make(map[string]*models.User)
	library.books[id] = newBook

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
