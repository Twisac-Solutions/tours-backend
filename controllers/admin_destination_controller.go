package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllDestinations(c *fiber.Ctx) error {
	destinations, err := services.GetAllDestinations()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve destinations"})
	}
	return c.JSON(destinations)
}

func GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	destination, err := services.GetDestinationByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Destination not found"})
	}
	return c.JSON(destination)
}

func CreateDestination(c *fiber.Ctx) error {
	var destination models.Destination
	if err := c.BodyParser(&destination); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, err := c.MultipartForm()
	if err == nil && form != nil {
		destination.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}

	err = services.CreateDestination(&destination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create destination"})
	}
	return c.JSON(destination)
}

func UpdateDestination(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.Destination
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	form, _ := c.MultipartForm()
	if form != nil {
		updated.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}

	err := services.UpdateDestination(id, &updated)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update destination"})
	}
	return c.JSON(updated)
}

func DeleteDestination(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteDestination(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete destination"})
	}
	return c.JSON(fiber.Map{"message": "Destination deleted"})
}
