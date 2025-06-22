# Go Bookstore API

A RESTful API for managing a bookstore built with Go, GORM, and Fiber.

## Project Architecture

This project follows clean architecture principles with a focus on:

- **Separation of Concerns**: Each package has a specific responsibility
- **Dependency Injection**: Components depend on abstractions, not concrete implementations
- **Testability**: The architecture allows for easy testing of each component

### Directory Structure

```
go-bookstore/
├── cmd/
│   ├── main.go          # Entry point of the application
│   └── server.go        # HTTP server configuration
├── pkg/
│   ├── config/          # Application configuration
│   │   └── db.go        # Database migration & setup
│   ├── handlers/        # HTTP request handlers
│   │   ├── book_handler.go    # Book API endpoints
│   │   └── health_handler.go  # Health check endpoint
│   ├── models/          # Domain models and business entities
│   │   └── book.go      # Book & Author models
│   ├── repository/      # Data access layer
│   │   ├── book.go         # Repository interfaces
│   │   └── impl/           # Repository implementations
│   │       └── book_repository.go
│   ├── service/         # Business logic layer
│   │   └── book_service.go   # Services that use repositories
│   └── utils/           # Utility functions
│       ├── env.go       # Environment variable helpers
│       ├── logger.go    # Logging utilities
│       └── must.go      # Error handling helpers
├── logs/                # Application logs
├── .dockerignore        # Docker ignore file
├── .env                 # Environment variables
├── .gitignore           # Git ignore file
├── Dockerfile           # Docker build configuration
├── docker-compose.yaml  # Docker Compose configuration
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── Makefile             # Build and run commands
```

## Architecture Layers

### 1. Models Layer
- Domain entities like `Book` and `Author`
- Pure data structures with validation tags

### 2. Repository Layer
- Provides data access interfaces and implementations
- GORM is used for database operations
- Each repository focuses on a specific domain model

### 3. Service Layer
- Business logic implementation
- Services use repositories for data access
- Implements validation, error handling, and domain rules

### 4. Handler Layer
- HTTP request handling
- Maps HTTP requests to service calls
- Formats responses in JSON

## API Endpoints

### Books API
- `GET /api/v1/books` - Get all books
- `GET /api/v1/books/:id` - Get book by ID
- `POST /api/v1/books/create` - Create a new book
- `PUT /api/v1/books/:id` - Update a book
- `DELETE /api/v1/books/:id` - Delete a book

### Health Check
- `GET /api/v1/health` - Check API health status

## Setup and Running

### Prerequisites
- Go 1.24+
- MySQL 8.0+
- Docker & Docker Compose (optional)

### Environment Variables
Copy the `.env.example` to `.env` and adjust the values:

```
# App configuration
ADDR="0.0.0.0"
PORT="8080"
API_VERSION="/api/v1"

# Database configuration
DB_USER="username"
DB_PASS="password"
DB_ADDR="localhost"
DB_PORT="3306"
DB_NAME="bookstore"
```

### Running with Makefile

```bash
# Build and run the application
make run

# Start the database only
make db

# Start the database in background
make db-d

# Stop the database
make db-stop
```

### Running with Docker Compose

```bash
# Start all services
docker compose up

# Start in background
docker compose up -d

# Stop services
docker compose down
```

## Development Principles

1. **API First**: API design comes before implementation
2. **Stateless**: No client state is stored on the server
3. **SOLID Principles**: Single responsibility, Open-closed, etc.
4. **Immutable Data**: Data objects are not modified after creation
5. **Dependency Injection**: All components use interfaces for dependencies
