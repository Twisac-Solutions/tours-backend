package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-username/tout-api/database"
	"github.com/your-username/tout-api/models"
)

func GetUserProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	var user models.User

	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(fiber.Map{"user": user})
}
