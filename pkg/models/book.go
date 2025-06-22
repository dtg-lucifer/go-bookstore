package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Book represents a book entity in the database
type Book struct {
	ID            string     `json:"id" gorm:"primaryKey;type:varchar(191);column:id;autoIncrement:false"`
	Name          string     `json:"name" validate:"required"`
	AuthorID      string     `json:"author_id" gorm:"type:varchar(191);column:author_id;not null"`
	Author        Author     `json:"author" validate:"required" gorm:"foreignKey:AuthorID;references:ID"`
	Publisher     string     `json:"publisher" validate:"required"`
	PublishedYear uint       `json:"published_year" validate:"required"`
	Description   string     `json:"description" validate:"required" gorm:"size:255"`
	Price         float64    `json:"price" validate:"required"`
	Pages         int        `json:"pages" validate:"required"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at" gorm:"index"`
}

// BeforeCreate is a GORM hook to generate UUID before creating a record
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = uuid.New().String()
	}
	return
}

// Author represents an author entity in the database
type Author struct {
	ID        string     `json:"id" gorm:"primaryKey;type:varchar(191);column:id;autoIncrement:false"`
	Name      string     `json:"name" validate:"required"`
	Bio       string     `json:"bio" validate:"required"`
	Books     []Book     `json:"books,omitempty" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}

// BeforeCreate is a GORM hook to generate UUID before creating a record
func (a *Author) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == "" {
		a.ID = uuid.New().String()
	}
	return
}
