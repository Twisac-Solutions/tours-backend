package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
)

func GetAllEvents() ([]models.Event, error) {
	var events []models.Event
	err := database.DB.Preload("Gallery").Preload("Schedule").Find(&events).Error
	return events, err
}

func GetEventByID(id string) (*models.Event, error) {
	var event models.Event
	err := database.DB.Preload("Gallery").Preload("Schedule").First(&event, "id = ?", id).Error
	return &event, err
}

func CreateEvent(event *models.Event) error {
	return database.DB.Create(event).Error
}

func UpdateEvent(id string, updated *models.Event) error {
	return database.DB.Model(&models.Event{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteEvent(id string) error {
	return database.DB.Delete(&models.Event{}, "id = ?", id).Error
}
