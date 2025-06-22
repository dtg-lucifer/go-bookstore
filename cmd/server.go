package main

import (
	"fmt"

	"github.com/dtg-lucifer/go-bookstore/pkg/config"
	"github.com/dtg-lucifer/go-bookstore/pkg/handlers"
	"github.com/dtg-lucifer/go-bookstore/pkg/repository/impl"
	"github.com/dtg-lucifer/go-bookstore/pkg/service"
	"github.com/dtg-lucifer/go-bookstore/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	App        *fiber.App
	Router     fiber.Router
	DB         *gorm.DB
	ApiVersion string
	Addr       string

	// Handlers
	bookHandler   *handlers.BookHandler
	healthHandler *handlers.HealthHandler
}

func NewServer(ip string, port string, version string) (*Server, error) {
	utils.Logger.Info("Initializing the Server")

	addr := fmt.Sprintf("%s:%s", ip, port)
	app := fiber.New(fiber.Config{
		AppName: "Book Store API",
	})

	if version == "" {
		return nil, fmt.Errorf("API version cannot be empty")
	}

	router := app.Group(version)

	return &Server{
		App:        app,
		Router:     router,
		Addr:       addr,
		ApiVersion: version,
		DB:         nil,
	}, nil
}

func (s *Server) SetupDB() error {
	utils.Logger.Info("Connecting to the Database")

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
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	utils.Logger.Info("Migrating the Database")
	config.MigrateDB(db)

	s.DB = db

	return nil
}

func (s *Server) SetupMiddlewares() error {
	utils.Logger.Info("Setting up Middlewares")

	writer, err := utils.NewEventLogger()
	if err != nil {
		return fmt.Errorf("failed to create event logger: %w", err)
	}

	s.App.Use(recover.New())
	s.App.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
	}))
	s.App.Use(requestid.New())
	s.App.Use(logger.New(logger.Config{
		Format:   "${green}[${time} - ${latency}]${reset} ${blue}[${ip}:${port}]${reset} ${blue}${locals:requestid}${reset} ${status} - [${method}] - ${yellow}${path}${reset} - ${blue}${queryParams} ${reqHeaders} ${body} - ${resBody}${reset}\n",
		TimeZone: "Asia/Kolkata",
	}))
	s.App.Use(logger.New(logger.Config{
		Format:   "${green}[${time} - ${latency}]${reset} ${blue}[${ip}:${port}]${reset} ${blue}${locals:requestid}${reset} ${status} - [${method}] - ${yellow}${path}${reset} - ${blue}${queryParams} ${reqHeaders} ${body} - ${resBody}${reset}\n",
		TimeZone: "Asia/Kolkata",
		Output:   writer,
	}))
	return nil
}

func (s *Server) SetupRoutes() error {
	utils.Logger.Info("Setting up Routes")

	// Initialize repositories
	bookRepo := impl.NewBookRepository(s.DB)

	// Initialize services
	bookService := service.NewBookService(bookRepo)

	// Initialize handlers
	s.bookHandler = handlers.NewBookHandler(bookService)
	s.healthHandler = handlers.NewHealthHandler()

	// Health routes
	s.Router.Get("/health", s.healthHandler.HealthCheck)

	// Book routes
	s.Router.Get("/books", s.bookHandler.GetAllBooks)
	s.Router.Get("/books/:id", s.bookHandler.GetBookById)
	s.Router.Post("/books/create", s.bookHandler.CreateBook)
	s.Router.Put("/books/:id", s.bookHandler.UpdateBook)
	s.Router.Delete("/books/:id", s.bookHandler.DeleteBook)

	return nil
}

func (s *Server) Start() error {
	return s.App.Listen(s.Addr)
}
