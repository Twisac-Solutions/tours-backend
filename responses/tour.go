package responses

import (
	"time"

	"github.com/Twisac-Solutions/tours-backend/models"
)

// TourResponse represents the API response format for a tour
type TourResponse struct {
	ID             string    `json:"id"`
	Title          string    `json:"title"`
	CategoryID     string    `json:"categoryId"`
	Description    string    `json:"description"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	PricePerPerson float64   `json:"pricePerPerson"`
	Currency       string    `json:"currency"`
	IsFeatured     bool      `json:"isFeatured"`
	CoverImage     string    `json:"coverImage"`
	Destination    struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"destination"`
	User struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		Username     string `json:"username"`
		ProfileImage string `json:"profileImage"`
		Role         string `json:"role"`
	} `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToTourResponse(tour models.Tour) TourResponse {
	response := TourResponse{
		ID:             tour.ID.String(),
		Title:          tour.Title,
		CategoryID:     tour.Category.String(),
		Description:    tour.Desc,
		StartDate:      tour.StartDate,
		EndDate:        tour.EndDate,
		PricePerPerson: tour.PricePerPerson,
		Currency:       tour.Currency,
		IsFeatured:     tour.IsFeatured,
		CreatedAt:      tour.CreatedAt,
		UpdatedAt:      tour.UpdatedAt,
	}

	// Set the cover image URL
	response.CoverImage = tour.CoverImage.URL

	// Set destination details
	response.Destination.ID = tour.Destination.ID.String()
	response.Destination.Name = tour.Destination.Name

	// Set user details
	response.User.ID = tour.User.ID.String()
	response.User.Name = tour.User.Name
	response.User.Username = tour.User.Username
	response.User.ProfileImage = "" // Set according to your User model structure
	response.User.Role = tour.User.Role

	return response
}
