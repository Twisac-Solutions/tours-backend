package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/gofiber/fiber/v2"
)

func GetAllReviews(c *fiber.Ctx) error {
	reviews, err := services.GetAllReviews()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve reviews"})
	}
	return c.JSON(reviews)
}

func GetReviewByID(c *fiber.Ctx) error {
	id := c.Params("id")
	review, err := services.GetReviewByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Review not found"})
	}
	return c.JSON(review)
}

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

func DeleteReview(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := services.DeleteReview(id); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete review"})
	}
	return c.JSON(fiber.Map{"message": "Review deleted"})
}
