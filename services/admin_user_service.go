package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/responses"
	"github.com/garrettladley/fiberpaginate/v2"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) ([]responses.UserResponse, int64, error) {
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
	if err != nil {
		return nil, 0, err
	}

	// Map DB models → response DTOs
	out := make([]responses.UserResponse, len(users))
	for i, u := range users {
		out[i] = responses.ToUserResponse(u)
	}
	return out, total, nil
}

func CreateUser(u *models.User) error {
	return database.DB.Create(u).Error
}

func UpdateUser(id string, updates *models.User) error {
	return database.DB.Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func DeleteUser(id string) error {
	return database.DB.Delete(&models.User{}, "id = ?", id).Error
}
