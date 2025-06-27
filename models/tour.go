package models

import (
	"time"

	"github.com/google/uuid"
)

type ItineraryItem struct {
	Day         int    `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Tour struct {
	ID    uuid.UUID `gorm:"type:text;primaryKey" json:"id"`
	Title string    `gorm:"type:varchar(255);not null" json:"title"`
	// Slug           string    `gorm:"uniqueIndex" json:"slug"`
	DestinationID uuid.UUID `json:"destinationId"`
	Category      uuid.UUID `json:"categoryId"`
	// ShortDesc     string    `json:"shortDescription"`
	Desc string `gorm:"type:text;not null" json:"fullDescription"`
	// DurationDays   int       `json:"durationDays"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
	PricePerPerson float64   `json:"pricePerPerson"`
	Currency       string    `json:"currency"`
	// GroupSize      int       `json:"groupSize"`
	// Availability   bool       `json:"availability"`
	IsFeatured bool `json:"isFeatured"`
	// Inclusions     []string  `gorm:"type:text[]" json:"inclusions"`
	// Exclusions     []string  `gorm:"type:text[]" json:"exclusions"`
	CoverImage MediaTour `gorm:"foreignKey:TourID" json:"coverImage"`
	// Gallery        []string  `gorm:"type:text[]" json:"gallery"`
	// Itinerary      []string  `gorm:"type:text[]" json:"itinerary"`
	// Tags           []string  `gorm:"type:text[]" json:"tags"`
	// Reviews        []string  `gorm:"type:text[]" json:"reviews"`
	CreatedBy uuid.UUID `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
