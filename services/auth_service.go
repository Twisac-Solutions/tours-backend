package services

import (
	"time"

	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	username := utils.GenerateUsername(data["name"])

	user := models.User{
		ID:              uuid.New(),
		Name:            data["name"],
		Email:           data["email"],
		Username:        username,
		Password:        string(hash),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		EmailVerifiedAt: time.Now(),
		IsVerified:      true,
	}

	database.DB.Create(&user)
	token, _ := utils.GenerateJWT(user.ID.String())
	return c.JSON(fiber.Map{"token": token})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.ID == uuid.Nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])) != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	token, _ := utils.GenerateJWT(user.ID.String())
	return c.JSON(fiber.Map{"token": token})
}

func GoogleSSO(c *fiber.Ctx) error {
	// Dummy Google login logic placeholder
	return c.JSON(fiber.Map{"message": "Google SSO not implemented"})
}
