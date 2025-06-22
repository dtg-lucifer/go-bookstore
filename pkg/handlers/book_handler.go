package handlers

import (
	"context"
	"net/http"

	"github.com/dtg-lucifer/go-bookstore/pkg/models"
	"github.com/dtg-lucifer/go-bookstore/pkg/service"
	"github.com/gofiber/fiber/v2"
)

// BookHandler handles HTTP requests related to books
type BookHandler struct {
	bookService service.BookService
}

// NewBookHandler creates a new BookHandler with the provided service
func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{
		bookService: service,
	}
}

// GetAllBooks handles GET /books request
func (h *BookHandler) GetAllBooks(ctx *fiber.Ctx) error {
	books, err := h.bookService.GetAllBooks(context.Background())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve books",
		})
	}

	if len(books) == 0 {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "No books found",
			"data":    []models.Book{},
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Books retrieved successfully",
		"data":    books,
	})
}

// GetBookById handles GET /books/:id request
func (h *BookHandler) GetBookById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Book ID is required",
		})
	}

	book, err := h.bookService.GetBookByID(context.Background(), id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book retrieved successfully",
		"data":    book,
	})
}

// CreateBook handles POST /books/create request
func (h *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	body := new(models.Book)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.bookService.CreateBook(context.Background(), body); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully",
		"data":    body,
	})
}

// UpdateBook handles PUT /books/:id request
func (h *BookHandler) UpdateBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Book ID is required",
		})
	}

	body := new(models.Book)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := h.bookService.UpdateBook(context.Background(), id, body); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Fetch the updated book to return in the response
	updatedBook, err := h.bookService.GetBookByID(context.Background(), id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Book updated but failed to retrieve the updated data",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book updated successfully",
		"data":    updatedBook,
	})
}

// DeleteBook handles DELETE /books/:id request
func (h *BookHandler) DeleteBook(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Book ID is required",
		})
	}

	if err := h.bookService.DeleteBook(context.Background(), id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}
