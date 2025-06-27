package models

import (
	"time"

	"github.com/google/uuid"
)

type Destination struct {
	ID          uuid.UUID        `gorm:"type:text;primaryKey;not null" json:"id"`
	Name        string           `gorm:"type:varchar(255);not null" json:"name"`
	Country     string           `gorm:"type:varchar(255);not null" json:"country"`
	Region      string           `gorm:"type:varchar(255);not null" json:"region"`
	Description string           `gorm:"type:text;not null" json:"description"`
	CoverImage  MediaDestination `gorm:"foreignKey:DestinationID" json:"coverImage"`
	// Gallery     []string  `gorm:"type:text[]" json:"gallery"`
	// Tours     []string  `gorm:"type:text[]" json:"tours"`
	// Events    []string  `gorm:"type:text[]" json:"events"`
	CreatedBy uuid.UUID `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
