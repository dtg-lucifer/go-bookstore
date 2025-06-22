package main

import (
	"os"

	"github.com/dtg-lucifer/go-bookstore/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize logger
	logger := utils.Logger

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Error("Failed to load .env file", "error", err)
		// Continue execution as default values will be used
	}

	// Create the server
	server, err := NewServer(
		utils.GetEnv("ADDR", "127.0.0.1"),
		utils.GetEnv("PORT", "8080"),
		utils.GetEnv("API_VERSION", "/api/vX"),
	)
	if err != nil {
		logger.Error("Failed to create server", "error", err)
		os.Exit(1)
	}

	// Set up database
	if err := server.SetupDB(); err != nil {
		logger.Error("Failed to set up database", "error", err)
		os.Exit(1)
	}

	// Set up middlewares
	if err := server.SetupMiddlewares(); err != nil {
		logger.Error("Failed to set up middlewares", "error", err)
		os.Exit(1)
	}

	// Set up routes
	if err := server.SetupRoutes(); err != nil {
		logger.Error("Failed to set up routes", "error", err)
		os.Exit(1)
	}

	// Start the server
	logger.Info("Starting server", "address", server.Addr)
	if err := server.Start(); err != nil {
		logger.Error("Server failed", "error", err)
		os.Exit(1)
	}
}
