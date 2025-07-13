package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/garrettladley/fiberpaginate/v2"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		pageInfo = &fiberpaginate.PageInfo{Page: 1, Limit: 10}
	}

	database.DB.Model(&models.User{}).Count(&total)

	err := database.DB.
		Order("created_at DESC").
		Offset(pageInfo.Start()).
		Limit(pageInfo.Limit).
		Find(&users).Error

	return users, total, err
}
