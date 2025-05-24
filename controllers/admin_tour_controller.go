package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllTours(c *fiber.Ctx) error {
	tours, err := services.GetAllTours()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve tours"})
	}
	return c.JSON(tours)
}

func GetTourByID(c *fiber.Ctx) error {
	id := c.Params("id")
	tour, err := services.GetTourByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tour not found"})
	}
	return c.JSON(tour)
}

func CreateTour(c *fiber.Ctx) error {
	var tour models.Tour
	if err := c.BodyParser(&tour); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, err := c.MultipartForm()
	if err == nil && form != nil {
		tour.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}

	err = services.CreateTour(&tour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create tour"})
	}
	return c.JSON(tour)
}

func UpdateTour(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.Tour
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid body"})
	}
	form, _ := c.MultipartForm()
	if form != nil {
		updated.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}
	err := services.UpdateTour(id, &updated)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update"})
	}
	return c.JSON(updated)
}

func DeleteTour(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteTour(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete"})
	}
	return c.JSON(fiber.Map{"message": "Tour deleted"})
}
