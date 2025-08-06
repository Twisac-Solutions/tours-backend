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

// CreateTourReview creates a review for a tour and updates the tour's rating
func CreateTourReview(review *models.Review) error {
	// Start a transaction
	tx := database.DB.Begin()

	// Create the review
	if err := tx.Create(review).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update tour rating
	if err := UpdateTourRating(review.TourID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// UpdateTourReview updates a review and recalculates the tour's rating
func UpdateTourReview(id string, updated *models.Review) error {
	// Start a transaction
	tx := database.DB.Begin()

	// Get the existing review to get the tour ID
	var existingReview models.Review
	if err := tx.First(&existingReview, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update the review
	if err := tx.Model(&models.Review{}).Where("id = ?", id).Updates(updated).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update tour rating
	if err := UpdateTourRating(existingReview.TourID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// DeleteTourReview deletes a review and updates the tour's rating
func DeleteTourReview(id string) error {
	// Start a transaction
	tx := database.DB.Begin()

	// Get the existing review to get the tour ID
	var existingReview models.Review
	if err := tx.First(&existingReview, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete the review
	if err := tx.Delete(&models.Review{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Update tour rating
	if err := UpdateTourRating(existingReview.TourID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// GetTourReviews returns all reviews for a specific tour
func GetTourReviews(tourID string) ([]models.Review, error) {
	var reviews []models.Review
	err := database.DB.Where("tour_id = ?", tourID).
		Preload("User").
		Order("created_at DESC").
		Find(&reviews).Error
	return reviews, err
}
