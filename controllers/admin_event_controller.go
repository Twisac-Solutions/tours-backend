package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

func GetAllEvents(c *fiber.Ctx) error {
	events, err := services.GetAllEvents()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve events"})
	}
	return c.JSON(events)
}

func GetEventByID(c *fiber.Ctx) error {
	id := c.Params("id")
	event, err := services.GetEventByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Event not found"})
	}
	return c.JSON(event)
}

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	form, err := c.MultipartForm()
	if err == nil && form != nil {
		event.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}
	err = services.CreateEvent(&event)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create event"})
	}
	return c.JSON(event)
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated models.Event
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	form, _ := c.MultipartForm()
	if form != nil {
		updated.CoverImage.URL, _ = utils.SaveFile(form.File["coverImage"])
	}
	err := services.UpdateEvent(id, &updated)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update event"})
	}
	return c.JSON(updated)
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteEvent(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete event"})
	}
	return c.JSON(fiber.Map{"message": "Event deleted"})
}
