package middlewares

import (
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func AdminProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, role, err := utils.VerifyJWTRole(c)
		if err != nil || role != "admin" {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
		}
		c.Locals("admin_id", userID)
		return c.Next()
	}
}
