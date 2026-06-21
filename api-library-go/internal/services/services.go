package services

import (
	"api-library-go/internal/models"
	"api-library-go/internal/storage"
)

func GetAllBooks() map[string]*models.Book {
	library := &storage.L

	return library.GetAllBooks()
}

func GetBook(id string) (*models.Book, error) {
	book, err := storage.L.GetBoook(id)
	return book, err
}

func CreateBook(book *models.Book) {
	storage.L.AddBook(book)
}

func UpBook(upBook *models.Book, id string) (*models.Book, error) {
	book, err := storage.L.UpBook(upBook, id)
	return book, err
}

func DelBook(id string) (*models.Book, error) {
	book, err := storage.L.DelBook(id)
	return book, err
}

// Users
func CreateUser(user *models.User) {
	storage.L.AddUser(user)
}

func GetAllUsers() map[string]*models.User {
	return storage.L.GetAllUsers()
}

func GetUser(id string) (*models.User, error) {
	user, err := storage.L.GetUser(id)
	return user, err
}

func UpUser(upUser *models.User, id string) (*models.User, error) {
	user, err := storage.L.UpUser(upUser, id)
	return user, err
}

func DelUser(id string) (*models.User, error) {
	user, err := storage.L.DelUser(id)
	return user, err
}
