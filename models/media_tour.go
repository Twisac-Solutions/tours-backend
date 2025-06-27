package models

import (
	"time"

	"github.com/google/uuid"
)

type MediaTour struct {
	ID        uint       `gorm:"primaryKey"`
	TourID    uuid.UUID  `gorm:"type:text;not null"`
	UserID    uuid.UUID  `gorm:"type:text;not null" json:"createdBy"`
	URL       string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index"`
}
