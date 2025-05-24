package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllEvents godoc
// @Summary      Get all events
// @Description  Retrieves a list of all events
// @Tags         events
// @Produce      json
// @Success      200  {array}   models.Event
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/events [get]
func GetAllEvents(c *fiber.Ctx) error {
	events, err := services.GetAllEvents()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve events"})
	}
	return c.JSON(events)
}

// GetEventByID godoc
// @Summary      Get event by ID
// @Description  Retrieves an event by its ID
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event ID"
// @Success      200  {object}  models.Event
// @Failure      404  {object}  models.ErrorResponse
// @Router       /admin/events/{id} [get]
func GetEventByID(c *fiber.Ctx) error {
	id := c.Params("id")
	event, err := services.GetEventByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Event not found"})
	}
	return c.JSON(event)
}

// CreateEvent godoc
// @Summary      Create a new event
// @Description  Creates a new event
// @Tags         events
// @Accept       multipart/form-data
// @Produce      json
// @Param        event  body      models.Event  true  "Event object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Event
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/events [post]
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

// UpdateEvent godoc
// @Summary      Update an event
// @Description  Updates an existing event by ID
// @Tags         events
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path      string      true  "Event ID"
// @Param        event body      models.Event true  "Event object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Event
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/events/{id} [put]
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

// DeleteEvent godoc
// @Summary      Delete an event
// @Description  Deletes an event by ID
// @Tags         events
// @Produce      json
// @Param        id   path      string  true  "Event ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/events/{id} [delete]
func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteEvent(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete event"})
	}
	return c.JSON(fiber.Map{"message": "Event deleted"})
}
