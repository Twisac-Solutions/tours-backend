package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
)

func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	err := database.DB.First(&category, "id = ?", id).Error
	return &category, err
}

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}

func UpdateCategory(id string, updated *models.Category) error {
	return database.DB.Model(&models.Category{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteCategory(id string) error {
	return database.DB.Delete(&models.Category{}, "id = ?", id).Error
}
