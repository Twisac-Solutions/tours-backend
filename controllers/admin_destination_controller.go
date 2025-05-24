package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllDestinations godoc
// @Summary      Get all destinations
// @Description  Retrieves a list of all destinations
// @Tags         destinations
// @Produce      json
// @Success      200  {array}   models.Destination
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/destinations [get]
func GetAllDestinations(c *fiber.Ctx) error {
	destinations, err := services.GetAllDestinations()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve destinations"})
	}
	return c.JSON(destinations)
}

// GetDestinationByID godoc
// @Summary      Get destination by ID
// @Description  Retrieves a destination by its ID
// @Tags         destinations
// @Produce      json
// @Param        id   path      string  true  "Destination ID"
// @Success      200  {object}  models.Destination
// @Failure      404  {object}  models.ErrorResponse
// @Router       /admin/destinations/{id} [get]
func GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	destination, err := services.GetDestinationByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Destination not found"})
	}
	return c.JSON(destination)
}

// CreateDestination godoc
// @Summary      Create a new destination
// @Description  Creates a new destination
// @Tags         destinations
// @Accept       multipart/form-data
// @Produce      json
// @Param        destination  body      models.Destination  true  "Destination object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Destination
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/destinations [post]
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

// UpdateDestination godoc
// @Summary      Update a destination
// @Description  Updates an existing destination by ID
// @Tags         destinations
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path      string      true  "Destination ID"
// @Param        destination body      models.Destination true  "Destination object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Destination
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/destinations/{id} [put]
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

// DeleteDestination godoc
// @Summary      Delete a destination
// @Description  Deletes a destination by ID
// @Tags         destinations
// @Produce      json
// @Param        id   path      string  true  "Destination ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/destinations/{id} [delete]
func DeleteDestination(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteDestination(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete destination"})
	}
	return c.JSON(fiber.Map{"message": "Destination deleted"})
}
