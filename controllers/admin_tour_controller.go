package controllers

import (
	"mime/multipart"
	"time"

	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateTourRequest struct {
	Title          string                `form:"title"`
	DestinationID  string                `form:"destinationId"`
	CategoryID     string                `form:"categoryId"`
	Description    string                `form:"desc"`
	StartDate      time.Time             `form:"startDate"`
	EndDate        time.Time             `form:"endDate"`
	PricePerPerson float64               `form:"pricePerPerson"`
	Currency       string                `form:"currency"`
	IsFeatured     bool                  `form:"isFeatured"`
	CoverImage     *multipart.FileHeader `form:"coverImage"`
}

type UpdateTourRequest struct {
	Title          string                `form:"title"`
	DestinationID  string                `form:"destinationId"`
	CategoryID     string                `form:"categoryId"`
	Description    string                `form:"desc"`
	StartDate      time.Time             `form:"startDate"`
	EndDate        time.Time             `form:"endDate"`
	PricePerPerson float64               `form:"pricePerPerson"`
	Currency       string                `form:"currency"`
	IsFeatured     bool                  `form:"isFeatured"`
	CoverImage     *multipart.FileHeader `form:"coverImage"`
}

// GetAllTours godoc
// @Summary      Get all tours
// @Description  Retrieves a list of all tours
// @Tags         admin_tours
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
// @Tags         admin_tours
// @Accept       multipart/form-data
// @Produce      json
// @Param        title          formData    string  true   "Tour title"
// @Param        destinationId  formData    string  true   "Destination ID"
// @Param        categoryId     formData    string  true   "Category ID"
// @Param        desc           formData    string  true   "Tour description"
// @Param        startDate      formData    string  true   "Start date"
// @Param        endDate        formData    string  true   "End date"
// @Param        pricePerPerson formData    number  true   "Price per person"
// @Param        currency       formData    string  true   "Currency"
// @Param        isFeatured     formData    boolean false  "Is featured"
// @Param        coverImage     formData    file    false  "Cover image"
// @Success      200  {object}  models.Tour
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/tours [post]
func CreateTour(c *fiber.Ctx) error {
	var req CreateTourRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Get user ID from context
	userID, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	tourID := uuid.New()
	tour := models.Tour{
		ID:             tourID,
		Title:          req.Title,
		DestinationID:  uuid.MustParse(req.DestinationID),
		Category:       uuid.MustParse(req.CategoryID),
		Desc:           req.Description,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		PricePerPerson: req.PricePerPerson,
		Currency:       req.Currency,
		IsFeatured:     req.IsFeatured,
		CreatedBy:      uuid.MustParse(userID),
	}

	// Handle cover image if provided
	if req.CoverImage != nil {
		fileURL, err := utils.SaveFile([]*multipart.FileHeader{req.CoverImage})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save cover image"})
		}
		tour.CoverImage = models.MediaTour{
			TourID:    tourID,
			UserID:    uuid.MustParse(userID),
			URL:       fileURL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	err := services.CreateTour(&tour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create tour"})
	}

	return c.JSON(tour)
}

// UpdateTour godoc
// @Summary      Update a tour
// @Description  Updates an existing tour
// @Tags         admin_tours
// @Accept       multipart/form-data
// @Produce      json
// @Param        id             path        string  true   "Tour ID"
// @Param        title          formData    string  true   "Tour title"
// @Param        destinationId  formData    string  true   "Destination ID"
// @Param        categoryId     formData    string  true   "Category ID"
// @Param        desc           formData    string  true   "Tour description"
// @Param        startDate      formData    string  true   "Start date"
// @Param        endDate        formData    string  true   "End date"
// @Param        pricePerPerson formData    number  true   "Price per person"
// @Param        currency       formData    string  true   "Currency"
// @Param        isFeatured     formData    boolean false  "Is featured"
// @Param        coverImage     formData    file    false  "Cover image"
// @Success      200  {object}  models.Tour
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/tours/{id} [put]
func UpdateTour(c *fiber.Ctx) error {
	id := c.Params("id")
	var req UpdateTourRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Get user ID from context
	userID, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found in context"})
	}

	// Get existing tour
	tour, err := services.GetTourByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tour not found"})
	}

	// Update tour fields
	tour.Title = req.Title
	tour.DestinationID = uuid.MustParse(req.DestinationID)
	tour.Category = uuid.MustParse(req.CategoryID)
	tour.Desc = req.Description
	tour.StartDate = req.StartDate
	tour.EndDate = req.EndDate
	tour.PricePerPerson = req.PricePerPerson
	tour.Currency = req.Currency
	tour.IsFeatured = req.IsFeatured

	// Handle cover image if provided
	if req.CoverImage != nil {
		fileURL, err := utils.SaveFile([]*multipart.FileHeader{req.CoverImage})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save cover image"})
		}
		tour.CoverImage = models.MediaTour{
			TourID:    uuid.MustParse(id),
			UserID:    uuid.MustParse(userID),
			URL:       fileURL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	err = services.UpdateTour(id, tour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update tour"})
	}

	return c.JSON(tour)
}

// DeleteTour godoc
// @Summary      Delete a tour
// @Description  Deletes a tour by ID
// @Tags         admin_tours
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
