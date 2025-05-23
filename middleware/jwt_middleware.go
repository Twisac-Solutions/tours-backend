package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-username/tout-api/utils"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, err := utils.VerifyJWT(c)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		c.Locals("userId", userId)
		return c.Next()
	}
}
