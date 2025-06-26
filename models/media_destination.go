package models

import (
	"time"

	"github.com/google/uuid"
)

type MediaType string

const (
	ImageType MediaType = "image"
	VideoType MediaType = "video"
	GifType   MediaType = "gif"
)

type MediaDestination struct {
	ID            uint       `gorm:"primaryKey"`
	DestinationID uuid.UUID  `gorm:"type:text;not null"`
	UserID        uuid.UUID  `gorm:"type:text;not null" json:"createdBy"`
	URL           string     `gorm:"type:varchar(255);not null"`
	Type          MediaType  `gorm:"type:varchar(20);not null"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `gorm:"index"`
}
