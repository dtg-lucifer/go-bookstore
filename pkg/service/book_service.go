package service

import (
	"context"
	"errors"

	"github.com/dtg-lucifer/go-bookstore/pkg/models"
	"github.com/dtg-lucifer/go-bookstore/pkg/repository"
)

// BookService defines the interface for book-related business logic
type BookService interface {
	GetAllBooks(ctx context.Context) ([]models.Book, error)
	GetBookByID(ctx context.Context, id string) (*models.Book, error)
	CreateBook(ctx context.Context, book *models.Book) error
	UpdateBook(ctx context.Context, id string, book *models.Book) error
	DeleteBook(ctx context.Context, id string) error
}

// BookServiceImpl implements the BookService interface
type BookServiceImpl struct {
	repo repository.BookRepository
}

// NewBookService creates a new BookService instance
func NewBookService(repo repository.BookRepository) BookService {
	return &BookServiceImpl{
		repo: repo,
	}
}

// GetAllBooks retrieves all books
func (s *BookServiceImpl) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	books, err := s.repo.GetAllBooks(ctx)
	if err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return []models.Book{}, nil
	}

	return books, nil
}

// GetBookByID retrieves a book by its ID
func (s *BookServiceImpl) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	if id == "" {
		return nil, errors.New("book ID cannot be empty")
	}

	return s.repo.GetBookByID(ctx, id)
}

// CreateBook creates a new book
func (s *BookServiceImpl) CreateBook(ctx context.Context, book *models.Book) error {
	if book == nil {
		return errors.New("book cannot be nil")
	}

	if book.Name == "" {
		return errors.New("book name cannot be empty")
	}

	if book.Author.Name == "" {
		return errors.New("author name cannot be empty")
	}

	return s.repo.CreateBook(ctx, book)
}

// UpdateBook updates an existing book
func (s *BookServiceImpl) UpdateBook(ctx context.Context, id string, book *models.Book) error {
	if id == "" {
		return errors.New("book ID cannot be empty")
	}

	if book == nil {
		return errors.New("book cannot be nil")
	}

	// First check if the book exists
	_, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.UpdateBook(ctx, id, book)
}

// DeleteBook deletes a book
func (s *BookServiceImpl) DeleteBook(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("book ID cannot be empty")
	}

	// First check if the book exists
	_, err := s.repo.GetBookByID(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.DeleteBook(ctx, id)
}
