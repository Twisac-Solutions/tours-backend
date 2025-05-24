package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
)

func GetAllReviews() ([]models.Review, error) {
	var reviews []models.Review
	err := database.DB.Find(&reviews).Error
	return reviews, err
}

func GetReviewByID(id string) (*models.Review, error) {
	var review models.Review
	err := database.DB.First(&review, "id = ?", id).Error
	return &review, err
}

func CreateReview(review *models.Review) error {
	return database.DB.Create(review).Error
}

func UpdateReview(id string, updated *models.Review) error {
	return database.DB.Model(&models.Review{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteReview(id string) error {
	return database.DB.Delete(&models.Review{}, "id = ?", id).Error
}
