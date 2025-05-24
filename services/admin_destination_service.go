package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
)

func GetAllDestinations() ([]models.Destination, error) {
	var destinations []models.Destination
	err := database.DB.Preload("Gallery").Find(&destinations).Error
	return destinations, err
}

func GetDestinationByID(id string) (*models.Destination, error) {
	var destination models.Destination
	err := database.DB.Preload("Gallery").First(&destination, "id = ?", id).Error
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
