package models

import (
	"time"
)

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

type Author struct {
	ID        string     `json:"id" gorm:"primaryKey;type:varchar(191);column:id;autoIncrement:false"`
	Name      string     `json:"name" validate:"required"`
	Bio       string     `json:"bio" validate:"required"`
	Books     []Book     `json:"books" gorm:"foreignKey:AuthorID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
