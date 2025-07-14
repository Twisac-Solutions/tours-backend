package controllers

import (
	"log"
	"time"

	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/requests"
	"github.com/Twisac-Solutions/tours-backend/responses"
	"github.com/Twisac-Solutions/tours-backend/services"
	"github.com/Twisac-Solutions/tours-backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetAllTours godoc
// @Summary      Get all tours
// @Description  Retrieves a list of all tours
// @Tags         tours
// @Produce      json
// @Param        page   query    integer  false  "Page number (default: 1)"
// @Param        limit  query    integer  false  "Limit per page (default: 10)"
// @Success      200  {object}   object{data=[]responses.TourResponse,meta=object{page=integer,limit=integer,total=integer,total_pages=integer}}
// @Failure      500  {object}  models.ErrorResponse
// @Router       /api/tours [get]
func GetAllTours(c *fiber.Ctx) error {
	tours, totalCount, err := services.GetAllTours(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve tours"})
	}

	// Convert tours to response format
	tourResponses := make([]responses.TourResponse, len(tours))
	for i, tour := range tours {
		tourResponses[i] = responses.ToTourResponse(tour)
	}

	return c.JSON(utils.PaginationResponse(c, tourResponses, totalCount))
}

// GetTourByID godoc
// @Summary      Get tour by ID
// @Description  Retrieves a tour by its ID
// @Tags         tours
// @Produce      json
// @Param        id   path      string  true  "Tour ID"
// @Success      200  {object}  responses.TourResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/tours/{id} [get]
func GetTourByID(c *fiber.Ctx) error {
	id := c.Params("id")
	tour, err := services.GetTourByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Tour not found"})
	}
	return c.JSON(responses.ToTourResponse(*tour))
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
// @Success      200  {object}  responses.TourResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/tours [post]
func CreateTour(c *fiber.Ctx) error {
	// Debug incoming request
	log.Println("Content-Type:", c.Get("Content-Type"))

	// Parse the multipart form directly to see if we can get the file
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing multipart form:", err)
	} else {
		log.Println("Form fields:", form.Value)
		log.Println("Form files:", form.File)
	}
	var req requests.CreateTourRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}
	validationResult := requests.ValidateCreateTourRequest(req)
	if !validationResult.Valid {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": validationResult.Errors,
		})
	}

	// Get user ID from context
	userID, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(500).JSON(fiber.Map{"error": "User ID not found in context"})
	}
	userUUID := uuid.MustParse(userID)
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user details"})
	}
	destination, err := services.GetDestinationByID(req.DestinationID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get Destination"})
	}

	tourID := uuid.New()
	tour := models.Tour{
		ID:             tourID,
		Title:          req.Title,
		DestinationID:  uuid.MustParse(req.DestinationID),
		Destination:    *destination,
		Category:       uuid.MustParse(req.CategoryID),
		Desc:           req.Description,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
		PricePerPerson: req.PricePerPerson,
		Currency:       req.Currency,
		IsFeatured:     req.IsFeatured,
		CreatedBy:      userUUID,
		User:           *user,
	}

	// Handle cover image separately from the request body parsing
	file, err := c.FormFile("coverImage")
	if err == nil && file != nil {
		log.Println("Cover Image Found manually:", file.Filename, file.Size)
		fileURL, err := utils.UploadImageToCloudinary(file, "tours")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to upload cover image to Cloudinary"})
		}
		tourID := uuid.New()
		tour.CoverImage = models.MediaTour{
			TourID:    tourID,
			UserID:    uuid.MustParse(userID),
			URL:       fileURL,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	} else if err != nil {
		log.Println("Error getting cover image:", err)
	}

	err = services.CreateTour(&tour)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create tour"})
	}
	createdTour, err := services.GetTourByID(tour.ID.String())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve created tour"})
	}

	return c.JSON(responses.ToTourResponse(*createdTour))
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
// @Success      200  {object}  responses.TourResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/tours/{id} [put]
func UpdateTour(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.UpdateTourRequest
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
	userUUID := uuid.MustParse(userID)
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user details"})
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
	tour.CreatedBy = userUUID
	tour.User = *user

	// Handle cover image if provided
	if req.CoverImage != nil {
		fileURL, err := utils.UploadImageToCloudinary(req.CoverImage, "tours")
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

	return c.JSON(responses.ToTourResponse(*tour))
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
