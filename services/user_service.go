package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/responses"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(c *fiber.Ctx) error {
	userId := c.Locals("userID").(string)
	var user models.User

	if err := database.DB.First(&user, "id = ?", userId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(fiber.Map{"user": &responses.UserResponse{
		ID:             user.ID.String(),
		Email:          user.Email,
		Name:           user.Name,
		Username:       user.Username,
		Role:           user.Role,
		ProfilePicture: user.ProfileImage.URL,
		Bio:            user.Bio,
		Phone:          user.Phone,
		Country:        user.Country,
		City:           user.City,
		Language:       user.Language,
		IsVerified:     user.IsVerified,
		CreatedAt:      user.CreatedAt,
	}})
}

func GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, "id = ?", id).Error
	return &user, err
}
