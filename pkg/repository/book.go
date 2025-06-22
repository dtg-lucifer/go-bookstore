package repository

import "github.com/dtg-lucifer/go-bookstore/pkg/models"

type BookRepository interface {
	GetAllBooks() []models.Book
	GetBookByID(id string) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(id string, book *models.Book) error
	DeleteBook(id string) error
}
