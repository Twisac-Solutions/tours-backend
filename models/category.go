package models

import (
	"time"

	"github.com/google/uuid"
)



type Category struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name"`
	Slug        string    `gorm:"uniqueIndex" json:"slug"`
	Description string    `json:"description"`
	Type        string    `json:"type"` // tour, event, both
	Icon        string    `json:"icon"` // emoji or UI icon
	CoverImage  Media     `gorm:"embedded" json:"coverImage"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
