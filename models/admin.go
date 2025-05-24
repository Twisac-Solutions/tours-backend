package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID        uuid.UUID `gorm:"type:text;primaryKey" json:"id"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	return
}
