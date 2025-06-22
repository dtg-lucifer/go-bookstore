package config

import (
	"github.com/dtg-lucifer/go-bookstore/pkg/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Book{},
		&models.Author{},
	)
	if err != nil {
		return err
	}

	return nil
}
