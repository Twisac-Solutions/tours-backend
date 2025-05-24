package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func AdminLogin(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	admin, err := services.FindAdminByEmail(input.Email)
	if err != nil || !utils.CheckPasswordHash(input.Password, admin.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, err := utils.GenerateJWTRole(admin.ID.String(), "admin")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token, "admin": admin})
}
