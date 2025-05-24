package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AdminOnly(c *fiber.Ctx) error {
	// Dummy example; replace with real role check
	role := c.Locals("userRole")
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}
	return c.Next()
}
