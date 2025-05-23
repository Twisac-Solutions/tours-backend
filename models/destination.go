package models

import (
	"time"

	"github.com/google/uuid"
)

type Destination struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string         `json:"name"`
	Slug        string         `gorm:"uniqueIndex" json:"slug"`
	Country     string         `json:"country"`
	Region      string         `json:"region"`
	Description string         `json:"description"`
	CoverImage  Media          `gorm:"embedded" json:"coverImage"`
	Gallery     pq.StringArray `gorm:"type:text[]" json:"gallery"`
	Tours       pq.StringArray `gorm:"type:text[]" json:"tours"`
	Events      pq.StringArray `gorm:"type:text[]" json:"events"`
	CreatedBy   uuid.UUID      `json:"createdBy"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
}
