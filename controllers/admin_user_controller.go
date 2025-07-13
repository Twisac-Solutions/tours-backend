package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
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

// GetAllUsers godoc
// @Summary      List all users (paginated)
// @Description  Returns paginated user summaries, newest first
// @Tags         admin_users
// @Produce      json
// @Param        page  query int false "Page number (default 1)"
// @Param        limit query int false "Items per page (default 10)"
// @Success      200  {object} object{data=[]responses.UserResponse,meta=object{page=int,limit=int,total=int,total_pages=int}}
// @Failure      500  {object} models.ErrorResponse
// @Router       /admin/users [get]
func GetAllUsers(c *fiber.Ctx) error {
	users, total, err := services.GetAllUsers(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(utils.PaginationResponse(c, users, total))
}
