package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MediaType string

const (
	ImageType MediaType = "image"
	VideoType MediaType = "video"
	GifType   MediaType = "gif"
)

type MediaDestination struct {
	gorm.Model
	DestinationID uuid.UUID `gorm:"type:text;not null"`
	UserID        uuid.UUID `gorm:"type:text;not null"`
	URL           string    `gorm:"type:varchar(255);not null"`
	Type          MediaType `gorm:"type:varchar(20);not null"`
}
