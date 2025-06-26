package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID `gorm:"type:text;primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"` // UI icon
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (m *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	return
}
