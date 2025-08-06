package services

import (
	"time"

	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/garrettladley/fiberpaginate/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

// GetFeaturedTours returns tours marked as featured (paginated)
func GetFeaturedTours(c *fiber.Ctx) ([]models.Tour, int64, error) {
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

	// Count total featured tours
	database.DB.Model(&models.Tour{}).Where("is_featured = ?", true).Count(&totalCount)

	// Get featured tours with pagination
	err := database.DB.Where("is_featured = ?", true).
		Offset(pageInfo.Start()).
		Limit(pageInfo.Limit).
		Preload("User").
		Preload("Destination").
		Preload("CoverImage").
		Order("created_at DESC").
		Find(&tours).Error

	return tours, totalCount, err
}

// func UpdateTour(id string, updated *models.Tour) error {
// 	return database.DB.Model(&models.Tour{}).Where("id = ?", id).Updates(map[string]interface{}{
// 		"title":            updated.Title,
// 		"destination_id":   updated.DestinationID,
// 		"category":         updated.Category,
// 		"description":      updated.Description,
// 		"about":            updated.About,
// 		"start_date":       updated.StartDate,
// 		"end_date":         updated.EndDate,
// 		"price_per_person": updated.PricePerPerson,
// 		"currency":         updated.Currency,
// 		"is_featured":      updated.IsFeatured,
// 	}).Error
// }

// GetFilteredTours returns tours based on various filter criteria
func GetFilteredTours(c *fiber.Ctx) ([]models.Tour, int64, error) {
	var tours []models.Tour
	var totalCount int64

	// Get pagination info from context
	pageInfo, ok := fiberpaginate.FromContext(c)
	if !ok {
		pageInfo = &fiberpaginate.PageInfo{
			Page:  1,
			Limit: 10,
		}
	}

	// Build query with filters
	query := database.DB.Model(&models.Tour{})

	// Filter by upcoming tours if requested
	if c.Query("upcoming") == "true" {
		query = query.Where("start_date > ?", time.Now())
	}

	// Filter by featured tours if requested
	if c.Query("featured") == "true" {
		query = query.Where("is_featured = ?", true)
	}

	// Filter by destination if provided
	if destinationID := c.Query("destination_id"); destinationID != "" {
		query = query.Where("destination_id = ?", destinationID)
	}

	// Filter by category if provided
	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category = ?", categoryID)
	}

	// Filter by price range if provided
	if minPrice := c.Query("min_price"); minPrice != "" {
		query = query.Where("price_per_person >= ?", minPrice)
	}
	if maxPrice := c.Query("max_price"); maxPrice != "" {
		query = query.Where("price_per_person <= ?", maxPrice)
	}

	// Count total records matching filters
	query.Count(&totalCount)

	// Get filtered tours with pagination
	err := query.Offset(pageInfo.Start()).
		Limit(pageInfo.Limit).
		Preload("User").
		Preload("Destination").
		Preload("CoverImage").
		Order("created_at DESC").
		Find(&tours).Error

	return tours, totalCount, err
}

// UpdateTourRating calculates and updates the average rating and review count for a tour
func UpdateTourRating(tourID uuid.UUID) error {
	var result struct {
		AverageRating float64
		ReviewCount   int64
	}

	// Calculate average rating and count from reviews
	err := database.DB.Model(&models.Review{}).
		Select("AVG(rating) as average_rating, COUNT(*) as review_count").
		Where("tour_id = ?", tourID).
		Scan(&result).Error

	if err != nil {
		return err
	}

	// Update the tour with the new rating information
	return database.DB.Model(&models.Tour{}).
		Where("id = ?", tourID).
		Updates(map[string]interface{}{
			"average_rating": result.AverageRating,
			"review_count":   result.ReviewCount,
		}).Error
}

// CreateTourWithReview creates a tour and optionally adds a review
func CreateTourWithReview(tour *models.Tour, review *models.Review) error {
	// Start a transaction
	tx := database.DB.Begin()

	// Create the tour
	if err := tx.Create(tour).Error; err != nil {
		tx.Rollback()
		return err
	}

	// If review is provided, create it
	if review != nil {
		review.TourID = tour.ID
		if err := tx.Create(review).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Update tour rating
		if err := UpdateTourRating(tour.ID); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
