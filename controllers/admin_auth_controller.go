package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AdminLoginResponse represents the admin login response payload.
type AdminLoginResponse struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Username       string `json:"username"`
	Role           string `json:"role"`
	ProfilePicture string `json:"profile_picture"`
}

// AdminLogin godoc
// @Summary      Admin login
// @Description  Authenticates an admin and returns a JWT token
// @Tags         admin_auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      AdminLoginRequest  true  "Admin login credentials"
// @Success      200  {object}  AdminLoginResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      401  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/login [post]
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

	return c.JSON(fiber.Map{"token": token, "user": &AdminLoginResponse{
		ID:       admin.ID.String(),
		Email:    admin.Email,
		Name:     admin.Name,
		Username: admin.Username,
		Role:     admin.Role,
	}})
}
