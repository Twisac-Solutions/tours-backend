package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/responses"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetCurrentAdminProfile godoc
// @Summary      Get current admin profile
// @Description  Returns the profile of the logged-in admin
// @Tags         admin_users
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

/* ---------- GET /admin/user/:id ---------- */
// GetUserByID godoc
// @Summary      Get user by id
// @Description  Returns one user
// @Tags         admin_users
// @Produce      json
// @Param        id  path  string  true  "User ID"
// @Success      200  {object} responses.UserResponse
// @Failure      404  {object} models.ErrorResponse
// @Router       /admin/user/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	user, err := services.GetUserByID(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(responses.ToUserResponse(*user))
}

/* ---------- POST /admin/user ---------- */
type createUserRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     string  `json:"name"  validate:"required"`
	Password string  `json:"password" validate:"required,min=6"`
	Role     string  `json:"role"  validate:"required,oneof=user admin superadmin"`
	Username *string `json:"username,omitempty"`
}

// CreateUser godoc
// @Summary      Create user
// @Description  Admin creates a new user
// @Tags         admin_users
// @Accept       json
// @Produce      json
// @Param        body  body  createUserRequest  true  "User data"
// @Success      201  {object} responses.UserResponse
// @Failure      400  {object} models.ErrorResponse
// @Failure      500  {object} models.ErrorResponse
// @Router       /admin/user [post]
func CreateUser(c *fiber.Ctx) error {
	var req createUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
	}

	user := models.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Name:     req.Name,
		Username: utils.GenerateUsername(req.Name),
		Password: utils.HashPassword(req.Password),
		Role:     req.Role,
	}
	if err := services.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(responses.ToUserResponse(user))
}

/* ---------- PUT /admin/user/:id ---------- */
type updateUserRequest struct {
	Email *string `json:"email,omitempty"` // pointer = optional
	Name  *string `json:"name,omitempty"`
	Role  *string `json:"role,omitempty" validate:"omitempty,oneof=user admin superadmin"`
}

// UpdateUser godoc
// @Summary      Update user
// @Description  Admin updates an existing user
// @Tags         admin_users
// @Accept       json
// @Produce      json
// @Param        id    path  string             true  "User ID"
// @Param        body  body  updateUserRequest  true  "Fields to update"
// @Success      200   {object} responses.UserResponse
// @Failure      400  {object} models.ErrorResponse
// @Failure      404  {object} models.ErrorResponse
// @Router       /admin/user/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	var req updateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
	}

	user, err := services.GetUserByID(c.Params("id"))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Role != nil {
		user.Role = *req.Role
	}

	if err := services.UpdateUser(user.ID.String(), user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(responses.ToUserResponse(*user))
}

/* ---------- DELETE /admin/user/:id ---------- */
// DeleteUser godoc
// @Summary      Delete user
// @Description  Admin deletes a user
// @Tags         admin_users
// @Param        id  path  string  true  "User ID"
// @Success      204 "No Content"
// @Failure      404  {object} models.ErrorResponse
// @Router       /admin/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	if err := services.DeleteUser(c.Params("id")); err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.SendStatus(204)
}
