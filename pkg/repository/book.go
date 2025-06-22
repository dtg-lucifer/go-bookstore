package repository

import (
	"context"

	"github.com/dtg-lucifer/go-bookstore/pkg/models"
)

// BookRepository defines the interface for book-related database operations
type BookRepository interface {
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) error
	UpdateBook(ctx context.Context, id string, book *models.Book) error
	DeleteBook(ctx context.Context, id string) error
}
