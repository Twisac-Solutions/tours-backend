package controllers

import (
	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
)

// GetAllTours godoc
// @Summary      Get all tours
// @Description  Retrieves a list of all tours
// @Tags         tours
// @Produce      json
// @Success      200  {array}   models.Tour
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/tours [get]
func GetAllTours(c *fiber.Ctx) error {
	tours, err := services.GetAllTours()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve tours"})
	}
	return c.JSON(tours)
}

// GetTourByID godoc
// @Summary      Get tour by ID
// @Description  Retrieves a tour by its ID
// @Tags         tours
// @Produce      json
// @Param        id   path      string  true  "Tour ID"
// @Success      200  {object}  models.Tour
// @Failure      404  {object}  models.ErrorResponse
// @Router       /admin/tours/{id} [get]
func GetTourByID(c *fiber.Ctx) error {
	id := c.Params("id")
	tour, err := services.GetTourByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tour not found"})
	}
	return c.JSON(tour)
}

// CreateTour godoc
// @Summary      Create a new tour
// @Description  Creates a new tour
// @Tags         tours
// @Accept       multipart/form-data
// @Produce      json
// @Param        tour  body      models.Tour  true  "Tour object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Tour
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/tours [post]
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

// UpdateTour godoc
// @Summary      Update a tour
// @Description  Updates an existing tour by ID
// @Tags         tours
// @Accept       multipart/form-data
// @Produce      json
// @Param        id    path      string      true  "Tour ID"
// @Param        tour  body      models.Tour true  "Tour object"
// @Param        coverImage formData file false "Cover image file"
// @Success      200   {object}  models.Tour
// @Failure      400   {object}  models.ErrorResponse
// @Failure      500   {object}  models.ErrorResponse
// @Router       /admin/tours/{id} [put]
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

// DeleteTour godoc
// @Summary      Delete a tour
// @Description  Deletes a tour by ID
// @Tags         tours
// @Produce      json
// @Param        id   path      string  true  "Tour ID"
// @Success      200  {object}  map[string]string
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/tours/{id} [delete]
func DeleteTour(c *fiber.Ctx) error {
	id := c.Params("id")
	err := services.DeleteTour(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete"})
	}
	return c.JSON(fiber.Map{"message": "Tour deleted"})
}
