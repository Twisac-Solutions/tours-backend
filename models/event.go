package models

import (
	"time"

	"github.com/google/uuid"
)

type ScheduleItem struct {
	Time        string `json:"time"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Event struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Title         string         `json:"title"`
	Slug          string         `gorm:"uniqueIndex" json:"slug"`
	DestinationID uuid.UUID      `json:"destinationId"`
	CategoryID    uuid.UUID      `json:"categoryId"`
	ShortDesc     string         `json:"shortDescription"`
	FullDesc      string         `json:"fullDescription"`
	EventDate     time.Time      `json:"eventDate"`
	DurationHours int            `json:"durationHours"`
	TicketPrice   float64        `json:"ticketPrice"`
	Currency      string         `json:"currency"`
	Capacity      int            `json:"capacity"`
	Availability  int            `json:"availability"`
	IsFeatured    bool           `json:"isFeatured"`
	Inclusions    pq.StringArray `gorm:"type:text[]" json:"inclusions"`
	Exclusions    pq.StringArray `gorm:"type:text[]" json:"exclusions"`
	CoverImage    Media          `gorm:"embedded" json:"coverImage"`
	Gallery       pq.StringArray `gorm:"type:text[]" json:"gallery"`
	Schedule      pq.StringArray `gorm:"type:text[]" json:"schedule"`
	Tags          pq.StringArray `gorm:"type:text[]" json:"tags"`
	Reviews       pq.StringArray `gorm:"type:text[]" json:"reviews"`
	CreatedBy     uuid.UUID      `json:"createdBy"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
}
