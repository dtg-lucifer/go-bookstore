package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/dtg-lucifer/go-bookstore/pkg/models"
	"github.com/dtg-lucifer/go-bookstore/pkg/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BookRepositoryImpl implements the BookRepository interface using GORM
type BookRepositoryImpl struct {
	DB *gorm.DB
}

// NewBookRepository creates a new BookRepository instance
func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &BookRepositoryImpl{
		DB: db,
	}
}

// GetAllBooks retrieves all books from the database
func (r *BookRepositoryImpl) GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	result := r.DB.WithContext(ctx).Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve books: %w", result.Error)
	}
	return books, nil
}

// GetBookByID retrieves a book by its ID
func (r *BookRepositoryImpl) GetBookByID(ctx context.Context, id string) (*models.Book, error) {
	var book models.Book
	result := r.DB.WithContext(ctx).Preload("Author").First(&book, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("book with ID %s not found", id)
		}
		return nil, fmt.Errorf("failed to retrieve book: %w", result.Error)
	}
	return &book, nil
}

// CreateBook creates a new book in the database
func (r *BookRepositoryImpl) CreateBook(ctx context.Context, book *models.Book) error {
	// Begin a transaction
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	// Generate UUIDs if they're empty
	if book.ID == "" {
		bookID, err := uuid.NewRandom()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to generate book UUID: %w", err)
		}
		book.ID = bookID.String()
	}

	// Check if we need to create a new author
	if book.Author.ID == "" {
		// Generate a new author ID
		authorID, err := uuid.NewRandom()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to generate author UUID: %w", err)
		}
		book.Author.ID = authorID.String()

		// Create the author
		if err := tx.Create(&book.Author).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create author: %w", err)
		}
	} else {
		// Author ID exists, check if we need to update author info
		var existingAuthor models.Author
		if err := tx.First(&existingAuthor, "id = ?", book.Author.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Author doesn't exist, create it
				if err := tx.Create(&book.Author).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to create author: %w", err)
				}
			} else {
				tx.Rollback()
				return fmt.Errorf("failed to check existing author: %w", err)
			}
		}
	}

	// Set the author ID in the book
	book.AuthorID = book.Author.ID

	// Create the book
	if err := tx.Create(book).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create book: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// UpdateBook updates an existing book in the database
func (r *BookRepositoryImpl) UpdateBook(ctx context.Context, id string, book *models.Book) error {
	// Begin a transaction
	tx := r.DB.WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	// Check if the book exists
	var existingBook models.Book
	if err := tx.First(&existingBook, "id = ?", id).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("book with ID %s not found", id)
		}
		return fmt.Errorf("failed to check existing book: %w", err)
	}

	// Update the author if needed
	if book.Author.ID != "" {
		// Check if the author exists
		var existingAuthor models.Author
		if err := tx.First(&existingAuthor, "id = ?", book.Author.ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// Create the author
				if err := tx.Create(&book.Author).Error; err != nil {
					tx.Rollback()
					return fmt.Errorf("failed to create author: %w", err)
				}
			} else {
				tx.Rollback()
				return fmt.Errorf("failed to check existing author: %w", err)
			}
		} else {
			// Update author fields if they exist
			existingAuthor.Name = book.Author.Name
			existingAuthor.Bio = book.Author.Bio
			if err := tx.Save(&existingAuthor).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to update author: %w", err)
			}
		}
	}

	// Update book fields
	book.ID = id // Ensure the ID is not changed
	if err := tx.Model(&existingBook).Updates(book).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update book: %w", err)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DeleteBook deletes a book from the database
func (r *BookRepositoryImpl) DeleteBook(ctx context.Context, id string) error {
	result := r.DB.WithContext(ctx).Delete(&models.Book{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete book: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("book with ID %s not found", id)
	}
	return nil
}
