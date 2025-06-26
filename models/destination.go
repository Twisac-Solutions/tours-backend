package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Destination struct {
	ID          uuid.UUID        `gorm:"type:text;primaryKey" json:"id"`
	Name        string           `json:"name"`
	Country     string           `json:"country"`
	Region      string           `json:"region"`
	Description string           `json:"description"`
	CoverImage  MediaDestination `gorm:"foreignKey:DestinationID" json:"coverImage"`
	// Gallery     []string  `gorm:"type:text[]" json:"gallery"`
	// Tours     []string  `gorm:"type:text[]" json:"tours"`
	// Events    []string  `gorm:"type:text[]" json:"events"`
	CreatedBy uuid.UUID `json:"createdBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *Destination) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}
