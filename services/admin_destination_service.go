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
	// Start a transaction
	tx := database.DB.Begin()

	// Update the destination
	if err := tx.Model(&models.Destination{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":        updated.Name,
		"description": updated.Description,
		"region":      updated.Region,
		"country":     updated.Country,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// If there's a new cover image, update it
	if updated.CoverImage.URL != "" {
		// First, delete the existing cover image
		if err := tx.Where("destination_id = ?", id).Delete(&models.MediaDestination{}).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Then create the new cover image
		updated.CoverImage.DestinationID = updated.ID
		if err := tx.Create(&updated.CoverImage).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func DeleteDestination(id string) error {
	return database.DB.Delete(&models.Destination{}, "id = ?", id).Error
}
