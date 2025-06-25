package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AdminOnly(c *fiber.Ctx) error {
	// Skip auth check for login route
	if c.Path() == "/admin/login" {
		return c.Next()
	}

	// Auth check for all other admin routes
	role := c.Locals("userRole")
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}
	return c.Next()
}
