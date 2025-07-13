package responses

import (
	"time"

	"github.com/Twisac-Solutions/tours-backend/models"
)

type UserResponse struct {
	ID             string    `json:"id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	ProfilePicture string    `json:"profile_picture"`
	Role           string    `json:"role"`
	Bio            string    `json:"bio"`
	Phone          string    `json:"phone"`
	Country        string    `json:"country"`
	City           string    `json:"city"`
	Language       string    `json:"language"`
	IsVerified     bool      `json:"is_verified"`
	CreatedAt      time.Time `json:"created_at"`
}

// Converts a models.User into the public response.
func ToUserResponse(u models.User) UserResponse {
	return UserResponse{
		ID:             u.ID.String(),
		Email:          u.Email,
		Name:           u.Name,
		Username:       u.Username,
		Role:           u.Role,
		ProfilePicture: u.ProfileImage.URL,
		Bio:            u.Bio,
		Phone:          u.Phone,
		Country:        u.Country,
		City:           u.City,
		Language:       u.Language,
		IsVerified:     u.IsVerified,
		CreatedAt:      u.CreatedAt,
	}
}
