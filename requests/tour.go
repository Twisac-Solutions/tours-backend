package requests

import (
	"fmt"
	"mime/multipart"
	"time"

	"github.com/Twisac-Solutions/tours-backend/utils"
)

type CreateTourRequest struct {
	Title          string                `json:"title" form:"title" required:"true"`
	DestinationID  string                `json:"destinationId" form:"destinationId" required:"true"`
	CategoryID     string                `json:"categoryId" form:"categoryId" required:"true"`
	Description    string                `json:"desc" form:"desc" required:"true"`
	CoverImage     *multipart.FileHeader `json:"coverImage" form:"coverImage"`
	StartDate      time.Time             `json:"startDate" form:"startDate" required:"true"`
	EndDate        time.Time             `json:"endDate" form:"endDate" required:"true"`
	PricePerPerson float64               `json:"pricePerPerson" form:"pricePerPerson" required:"true" min:"0"`
	Currency       string                `json:"currency" form:"currency" required:"true"`
	IsFeatured     bool                  `json:"isFeatured" form:"isFeatured"`
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

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateCreateTourRequest(req CreateTourRequest) utils.ValidationResult {
	validator := utils.NewValidator()

	// Add custom validators
	validator.AddCustomValidator("validateDates", func(i interface{}) error {
		req, ok := i.(CreateTourRequest)
		if !ok {
			return fmt.Errorf("invalid type for date validation")
		}
		if req.EndDate.Before(req.StartDate) {
			return fmt.Errorf("end date must be after start date")
		}
		return nil
	})

	return validator.Validate(req)
}
