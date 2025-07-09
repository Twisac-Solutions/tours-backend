package responses

import (
	"time"

	"github.com/Twisac-Solutions/tours-backend/models"
)

type DestinationResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Region      string `json:"region"`
	Country     string `json:"country"`
	CoverImage  string `json:"coverImage"`
	User        struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToDestinationResponse(destination models.Destination) DestinationResponse {
	response := DestinationResponse{
		ID:          destination.ID.String(),
		Name:        destination.Name,
		Description: destination.Description,
		Region:      destination.Region,
		Country:     destination.Country,
		CreatedAt:   destination.CreatedAt,
		UpdatedAt:   destination.UpdatedAt,
	}

	response.CoverImage = destination.CoverImage.URL

	response.User.ID = destination.User.ID.String()
	response.User.Name = destination.User.Name

	return response
}
