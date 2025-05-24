package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/gofiber/fiber/v2"
)

// GetAllReviews godoc
// @Summary      Get all reviews
// @Description  Retrieves a list of all reviews
// @Tags         admin_reviews
// @Produce      json
// @Success      200  {array}   models.Review
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/reviews [get]
func GetAllReviews(c *fiber.Ctx) error {
	reviews, err := services.GetAllReviews()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve reviews"})
	}
	return c.JSON(reviews)
}

// GetReviewByID godoc
// @Summary      Get review by ID
// @Description  Retrieves a review by its ID
// @Tags         admin_reviews
// @Produce      json
// @Param        id   path      string  true  "Review ID"
// @Success      200  {object}  models.Review
// @Failure      404  {object}  models.ErrorResponse
// @Router       /admin/reviews/{id} [get]
func GetReviewByID(c *fiber.Ctx) error {
	id := c.Params("id")
	review, err := services.GetReviewByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Review not found"})
	}
	return c.JSON(review)
}

// CreateReview godoc
// @Summary      Create a new review
// @Description  Creates a new review
// @Tags         admin_reviews
// @Accept       json
// @Produce      json
// @Param        review  body      models.Review  true  "Review object"
// @Success      200   {object}  models.Review
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/reviews [post]
func CreateReview(c *fiber.Ctx) error {
	var review models.Review
	if err := c.BodyParser(&review); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := services.CreateReview(&review); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create review"})
	}
	return c.JSON(review)
}

// UpdateReview godoc
// @Summary      Update a review
// @Description  Updates an existing review by ID
// @Tags         admin_reviews
// @Accept       json
// @Produce      json
// @Param        id    path      string      true  "Review ID"
// @Param        review body      models.Review true  "Review object"
// @Success      200   {object}  models.Review
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/reviews/{id} [put]
func UpdateReview(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.Review
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := services.UpdateReview(id, &updated); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update review"})
	}
	return c.JSON(updated)
}

// DeleteReview godoc
// @Summary      Delete a review
// @Description  Deletes a review by ID
// @Tags         admin_reviews
// @Produce      json
// @Param        id   path      string  true  "Review ID"
// @Success      200  {object}  models.MessageResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/reviews/{id} [delete]
func DeleteReview(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteReview(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete review"})
	}
	return c.JSON(fiber.Map{"message": "Review deleted"})
}
