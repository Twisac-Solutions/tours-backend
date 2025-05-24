package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

// UpdateAdminPassword allows the current admin to update their password.
func UpdateAdminPassword(c *fiber.Ctx) error {
	type Request struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}

	var body Request
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	admin := c.Locals("admin").(*models.User)

	if !utils.CheckPasswordHash(body.OldPassword, admin.Password) {
		return fiber.NewError(fiber.StatusUnauthorized, "Old password is incorrect")
	}

	hashed := utils.HashPassword(body.NewPassword)
	admin.Password = hashed

	database.DB.Save(&admin)

	return c.JSON(fiber.Map{"message": "Password updated"})
}

// ListAdmins godoc
// @Summary      List all admins
// @Description  Lists all admin users (super admin only)
// @Tags         admins
// @Produce      json
// @Success      200  {array}   models.User
// @Router       /admin [get]
func ListAdmins(c *fiber.Ctx) error {
	var admins []models.User
	database.DB.Where("role = ?", "admin").Find(&admins)
	return c.JSON(admins)
}

// CreateAdmin godoc
// @Summary      Create admin
// @Description  Creates a new admin user (super admin only)
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        body  body      models.User  true  "Admin user object"
// @Success      200   {object}  models.User
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin [post]
func CreateAdmin(c *fiber.Ctx) error {
	var data models.User
	if err := c.BodyParser(&data); err != nil {
		return fiber.ErrBadRequest
	}
	data.ID = utils.GenerateUUID()
	data.Role = "admin"
	data.Password = utils.HashPassword(data.Password)

	if err := database.DB.Create(&data).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(data)
}

// UpdateAdmin godoc
// @Summary      Update admin
// @Description  Updates an admin's name or email (super admin only)
// @Tags         admins
// @Accept       json
// @Produce      json
// @Param        id    path      string  true  "Admin ID"
// @Param        body  body      UpdateAdminRequest  true  "Fields to update"
// @Success      200   {object}  models.User
// @Failure      400   {object}  models.ErrorResponse
// @Failure      404   {object}  models.ErrorResponse
// @Router       /admin/{id} [put]
func UpdateAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.User
	if err := database.DB.First(&data, "id = ? AND role = ?", id, "admin").Error; err != nil {
		return fiber.ErrNotFound
	}

	var body map[string]string
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	if name, ok := body["name"]; ok {
		data.Name = name
	}
	if email, ok := body["email"]; ok {
		data.Email = email
	}

	database.DB.Save(&data)
	return c.JSON(data)
}

// DeleteAdmin godoc
// @Summary      Delete admin
// @Description  Deletes an admin user (super admin only)
// @Tags         admins
// @Produce      json
// @Param        id   path      string  true  "Admin ID"
// @Success      204  "No Content"
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/{id} [delete]
func DeleteAdmin(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Where("id = ? AND role = ?", id, "admin").Delete(&models.User{}).Error; err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}
