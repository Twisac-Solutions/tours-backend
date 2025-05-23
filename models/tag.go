package models

import (
	"time"

	"github.com/google/uuid"
)

type Tag struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `json:"name"`
	Slug        string    `gorm:"uniqueIndex" json:"slug"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Icon        string    `json:"icon"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
