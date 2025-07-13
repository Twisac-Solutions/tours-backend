package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/gofiber/fiber/v2"
)

// GetCurrentAdminProfile godoc
// @Summary      Get current admin profile
// @Description  Returns the profile of the logged-in admin
// @Tags         admin_user
// @Produce      json
// @Success      200  {object} models.User
// @Failure      401  {object} models.ErrorResponse
// @Router       /admin/user/me [get]
func GetCurrentAdminProfile(c *fiber.Ctx) error {
	return services.GetUserProfile(c)
}
