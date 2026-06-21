package services

import (
	"api-library-go/internal/models"
	"api-library-go/internal/storage"
)

func GetAllBooks() map[string]*models.Book {
	library := &storage.L

	return library.GetAllBooks()
}

func CreateBook(book *models.Book) {
	storage.L.AddBook(book)
}

func UpBook(upBook *models.Book, id string) (*models.Book, error) {
	book, err := storage.L.UpBook(upBook, id)
	return book, err
}
