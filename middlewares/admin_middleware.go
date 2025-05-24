package middlewares

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId, role, err := utils.VerifyJWTRole(c)
		if err != nil || (role != "admin" && role != "superadmin") {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}

		var user models.User
		if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Admin not found")
		}
		c.Locals("admin", &user)
		return c.Next()
	}
}

func SuperAdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		admin := c.Locals("admin").(*models.User)
		if admin.Role != "superadmin" {
			return fiber.NewError(fiber.StatusForbidden, "Super Admins only")
		}
		return c.Next()
	}
}
