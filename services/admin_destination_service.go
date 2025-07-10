package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/garrettladley/fiberpaginate/v2"
	"github.com/gofiber/fiber/v2"
)

func GetAllDestinations(c *fiber.Ctx) ([]models.Destination, int64, error) {
	var destinations []models.Destination
	var totalCount int64

	// Get pagination info from context
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		// If pagination info is not available, use default values
		pageInfo = &fiberpaginate.PageInfo{
			Page:  1,
			Limit: 10,
		}
	}

	// Count total records
	database.DB.Model(&models.Destination{}).Count(&totalCount)
	err := database.DB.Offset(pageInfo.Start()).
		Limit(pageInfo.Limit).Preload("User").Preload("CoverImage").Find(&destinations).Error
	return destinations, totalCount, err
}

func GetDestinationByID(id string) (*models.Destination, error) {
	var destination models.Destination
	err := database.DB.Preload("User").Preload("CoverImage").First(&destination, "id = ?", id).Error
	return &destination, err
}

func CreateDestination(destination *models.Destination) error {
	return database.DB.Create(destination).Error
}

func UpdateDestination(id string, updated *models.Destination) error {
	return database.DB.Model(&models.Destination{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteDestination(id string) error {
	return database.DB.Delete(&models.Destination{}, "id = ?", id).Error
}
