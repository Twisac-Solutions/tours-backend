package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
)

func GetAllTours() ([]models.Tour, error) {
	var tours []models.Tour
	// err := database.DB.Preload("Gallery").Preload("Itinerary").Find(&tours).Error
	err := database.DB.Find(&tours).Error
	return tours, err
}

func GetTourByID(id string) (*models.Tour, error) {
	var tour models.Tour
	err := database.DB.First(&tour, "id = ?", id).Error
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
