package main

import (
	"fmt"

	"github.com/dtg-lucifer/go-bookstore/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	App        *fiber.App
	DB         *gorm.DB
	ApiVersion string
	Addr       string
}

func NewServer(ip string, port string, version string) (*Server, error) {
	addr := fmt.Sprintf("%s:%s", ip, port)
	app := fiber.New(fiber.Config{
		AppName: "Book Store API",
	})

	if version == "" {
		return nil, fmt.Errorf("API version cannot be empty")
	}

	db_user := utils.GetEnv("DB_USER", "demo")
	db_pass := utils.GetEnv("DB_PASS", "password")
	db_addr := utils.GetEnv("DB_ADDR", "127.0.0.1")
	db_port := utils.GetEnv("DB_PORT", "3306")
	db_name := utils.GetEnv("DB_NAME", "book_store")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		db_user,
		db_pass,
		db_addr,
		db_port,
		db_name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return &Server{
		App:        app,
		Addr:       addr,
		ApiVersion: version,
		DB:         db,
	}, nil
}

func (s *Server) SetupRoutes() error {

	return nil
}

func (s *Server) Start() error {

	return nil
}
