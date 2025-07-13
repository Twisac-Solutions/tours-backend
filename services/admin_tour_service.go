package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/garrettladley/fiberpaginate/v2"
	"github.com/gofiber/fiber/v2"
)

func GetAllTours(c *fiber.Ctx) ([]models.Tour, int64, error) {
	var tours []models.Tour
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
	database.DB.Model(&models.Tour{}).Count(&totalCount)
	// err := database.DB.Preload("Gallery").Preload("Itinerary").Find(&tours).Error
	err := database.DB.Offset(pageInfo.Start()).
		Limit(pageInfo.Limit).Preload("User").Preload("Destination").Preload("CoverImage").Order("created_at DESC").Find(&tours).Error
	return tours, totalCount, err
}

func GetTourByID(id string) (*models.Tour, error) {
	var tour models.Tour
	err := database.DB.Preload("User").Preload("Destination").Preload("CoverImage").First(&tour, "id = ?", id).Error
	return &tour, err
}

func CreateTour(tour *models.Tour) error {
	return database.DB.Create(tour).Error
}

func UpdateTour(id string, updated *models.Tour) error {
	return database.DB.Model(&models.Tour{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteTour(id string) error {
	return database.DB.Delete(&models.Tour{}, "id = ?", id).Error
}
