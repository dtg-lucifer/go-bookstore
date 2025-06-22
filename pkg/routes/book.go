package routes

import (
	"fmt"

	"github.com/dtg-lucifer/go-bookstore/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bookhandler struct {
	DB *gorm.DB
}

func (h *Bookhandler) CreateBook(ctx *fiber.Ctx) error {
	body := new(models.Book)
	if err := ctx.BodyParser(body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	fmt.Printf("Received book data: %#v\n", body)

	bookId, err := uuid.NewRandom()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate UUID",
		})
	}

	authorId, err := uuid.NewRandom()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate UUID",
		})
	}

	body.ID = bookId.String()
	body.Author.ID = authorId.String()

	if err := h.DB.Create(&body.Author).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create author",
		})
	}

	tx := h.DB.Create(body)
	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to create book: %v", tx.Error),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Book created successfully",
		"data":    body,
	})

}

func (h *Bookhandler) GetAllBooks(ctx *fiber.Ctx) error {
	var books []models.Book
	tx := h.DB.Preload("Author").Find(&books)

	if len(books) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No books found",
			"data":    []any{},
		})
	}

	if tx.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve books",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Books retrieved successfully",
		"data":    books,
	})
}
