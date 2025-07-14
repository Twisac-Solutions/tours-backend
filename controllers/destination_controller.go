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

// GetAllDestinations godoc
// @Summary      Get all destinations
// @Description  Retrieves a list of all destinations
// @Tags         destinations
// @Produce      json
// @Param        page   query    integer  false  "Page number (default: 1)"
// @Param        limit  query    integer  false  "Limit per page (default: 10)"
// @Success      200  {object}   object{data=[]responses.DestinationResponse,meta=object{page=integer,limit=integer,total=integer,total_pages=integer}}
// @Failure      500  {object}  models.ErrorResponse
// @Router       /api/destinations [get]
func GetAllDestinations(c *fiber.Ctx) error {
	destinations, totalCount, err := services.GetAllDestinations(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve destinations"})
	}

	response := make([]responses.DestinationResponse, len(destinations))
	for i, d := range destinations {
		response[i] = responses.ToDestinationResponse(d)
	}

	return c.JSON(utils.PaginationResponse(c, response, totalCount))
}

// GetDestinationByID godoc
// @Summary      Get destination by ID
// @Description  Retrieves a destination by its ID
// @Tags         destinations
// @Produce      json
// @Param        id   path      string  true  "Destination ID"
// @Success      200  {object}  responses.DestinationResponse
// @Failure      404  {object}  models.ErrorResponse
// @Router       /api/destinations/{id} [get]
func GetDestinationByID(c *fiber.Ctx) error {
	id := c.Params("id")
	destination, err := services.GetDestinationByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Destination not found"})
	}
	return c.JSON(responses.ToDestinationResponse(*destination))
}

// CreateDestination godoc
// @Summary      Create a new destination
// @Description  Creates a new destination with optional cover image
// @Tags         admin_destinations
// @Accept       multipart/form-data
// @Produce      json
// @Param        name        formData    string  true   "Destination name"
// @Param        description formData    string  true   "Destination description"
// @Param        region      formData    string  true   "Destination region"
// @Param        country     formData    string  true   "Destination country"
// @Param        coverImage  formData    file    false  "Cover image file"
// @Success      200  {object}  responses.DestinationResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/destinations [post]
func CreateDestination(c *fiber.Ctx) error {
	var req requests.CreateDestinationRequest
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

	destinationID := uuid.New()
	userUUID := uuid.MustParse(userID)
	// Get the user
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user details"})
	}
	destination := models.Destination{
		ID:          destinationID,
		Name:        req.Name,
		Description: req.Description,
		Region:      req.Region,
		Country:     req.Country,
		CreatedBy:   userUUID,
		User:        *user,
	}

	// Handle cover image if provided
	file, err := c.FormFile("coverImage")
	if err == nil && file != nil {
		log.Println("Destination Cover Image Found manually:", file.Filename, file.Size)
		fileURL, err := utils.UploadImageToCloudinary(file, "destinations")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save cover image"})
		}
		destination.CoverImage = models.MediaDestination{
			ID:            uint(time.Now().Unix()), // or use auto-increment
			DestinationID: destinationID,
			UserID:        userUUID,
			URL:           fileURL,
			Type:          destination.CoverImage.Type,
		}
	}
	err = services.CreateDestination(&destination) // Changed from := to =
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create destination"})
	}

	createdDestination, err := services.GetDestinationByID(destination.ID.String())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve created destination"})
	}

	return c.JSON(responses.ToDestinationResponse(*createdDestination))
}

// UpdateDestination godoc
// @Summary      Update a destination
// @Description  Updates an existing destination
// @Tags         admin_destinations
// @Accept       multipart/form-data
// @Produce      json
// @Param        id          path        string  true   "Destination ID"
// @Param        name        formData    string  true   "Destination name"
// @Param        description formData    string  true   "Destination description"
// @Param        region      formData    string  true   "Destination region"
// @Param        country     formData    string  true   "Destination country"
// @Param        coverImage  formData    file    false  "Cover image file"
// @Success      200  {object}  responses.DestinationResponse
// @Failure      400  {object}  models.ErrorResponse
// @Failure      404  {object}  models.ErrorResponse
// @Failure      500  {object}  models.ErrorResponse
// @Router       /admin/destinations/{id} [put]
func UpdateDestination(c *fiber.Ctx) error {
	id := c.Params("id")
	var req requests.UpdateDestinationRequest
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

	// Get existing destination
	destination, err := services.GetDestinationByID(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Destination not found"})
	}

	// Get the user
	user, err := services.GetUserByID(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get user details"})
	}

	// Update fields
	destination.Name = req.Name
	destination.Description = req.Description
	destination.Region = req.Region
	destination.Country = req.Country
	destination.User = *user // Update the User relationship

	// Handle cover image if provided
	file, err := c.FormFile("coverImage")
	if err == nil && file != nil {
		log.Println("Destination Cover Image Found manually:", file.Filename, file.Size)
		fileURL, err := utils.UploadImageToCloudinary(file, "destinations")
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to save cover image"})
		}
		destination.CoverImage = models.MediaDestination{
			ID:            uint(time.Now().Unix()),
			DestinationID: destination.ID,
			UserID:        uuid.MustParse(userID),
			URL:           fileURL,
			Type:          destination.CoverImage.Type,
		}
	}

	err = services.UpdateDestination(id, destination)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update destination"})
	}

	return c.JSON(responses.ToDestinationResponse(*destination))
}

// DeleteDestination godoc
// @Summary      Delete a destination
// @Description  Deletes a destination by ID
// @Tags         admin_destinations
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
