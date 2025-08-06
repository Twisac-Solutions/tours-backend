package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey" json:"id"`
	UserID    uuid.UUID `json:"userId"`
	TourID    uuid.UUID `json:"tourId"`                        // Changed from PackageID to TourID for clarity
	Rating    int       `json:"rating" validate:"min=1,max=5"` // 1 to 5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`

	// Relationships
	User User `gorm:"foreignKey:UserID" json:"user"`
	Tour Tour `gorm:"foreignKey:TourID" json:"tour"`
}
