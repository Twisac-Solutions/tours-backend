package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey" json:"id"`
	UserID    uuid.UUID `json:"userId"`
	PackageID uuid.UUID `json:"packageId"` // Refers to either a Tour or an Event
	Rating    int       `json:"rating"`    // 1 to 5
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
}
