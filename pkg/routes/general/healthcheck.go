package general

import "github.com/gofiber/fiber/v2"

func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"status":  "ok",
		"message": "Service is running",
	})
}
