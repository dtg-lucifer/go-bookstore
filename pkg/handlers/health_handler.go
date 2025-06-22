package handlers

import "github.com/gofiber/fiber/v2"

// HealthHandler handles health check requests
type HealthHandler struct{}

// NewHealthHandler creates a new HealthHandler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck handles GET /health request
func (h *HealthHandler) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"message": "Service is running",
	})
}
