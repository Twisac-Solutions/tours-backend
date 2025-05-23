package models

import (
	"time"

	"github.com/google/uuid"
)

type Media struct {
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type ItineraryItem struct {
	Day         int    `json:"day"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Tour struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title          string         `json:"title"`
	Slug           string         `gorm:"uniqueIndex" json:"slug"`
	DestinationID  uuid.UUID      `json:"destinationId"`
	CategoryID     uuid.UUID      `json:"categoryId"`
	ShortDesc      string         `json:"shortDescription"`
	FullDesc       string         `json:"fullDescription"`
	DurationDays   int            `json:"durationDays"`
	StartDate      time.Time      `json:"startDate"`
	EndDate        time.Time      `json:"endDate"`
	PricePerPerson float64        `json:"pricePerPerson"`
	Currency       string         `json:"currency"`
	GroupSize      int            `json:"groupSize"`
	Availability   int            `json:"availability"`
	IsFeatured     bool           `json:"isFeatured"`
	Inclusions     pq.StringArray `gorm:"type:text[]" json:"inclusions"`
	Exclusions     pq.StringArray `gorm:"type:text[]" json:"exclusions"`
	CoverImage     Media          `gorm:"embedded" json:"coverImage"`
	Gallery        pq.StringArray `gorm:"type:text[]" json:"gallery"`
	Itinerary      pq.StringArray `gorm:"type:text[]" json:"itinerary"`
	Tags           pq.StringArray `gorm:"type:text[]" json:"tags"`
	Reviews        pq.StringArray `gorm:"type:text[]" json:"reviews"`
	CreatedBy      uuid.UUID      `json:"createdBy"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}
