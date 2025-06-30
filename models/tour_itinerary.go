package models

type Itinerary struct {
	Day         int    `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
