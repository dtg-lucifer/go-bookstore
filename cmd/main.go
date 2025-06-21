package main

import (
	"fmt"

	"github.com/dtg-lucifer/go-bookstore/pkg/utils"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	server := utils.Must(
		NewServer(
			utils.GetEnv("ADDR", "127.0.0.1"),
			utils.GetEnv("PORT", "8080"),
			utils.GetEnv("API_VERSION", "/api/vX"),
		),
	)

	if err := server.SetupRoutes(); err != nil {
		fmt.Printf("Error setting up routes: %v\n", err)
		return
	}

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}

	fmt.Println("Hello, world")
}
