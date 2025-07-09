package responses

import "time"

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
