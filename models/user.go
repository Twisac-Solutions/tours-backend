package models

import (
	"time"

	"github.com/google/uuid"
)

type ProfileImage struct {
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type SocialLinks struct {
	Instagram string `json:"instagram"`
	Facebook  string `json:"facebook"`
	Website   string `json:"website"`
}

type User struct {
	ID              uuid.UUID    `gorm:"type:text;primaryKey" json:"id"`
	Name            string       `json:"name"`
	Username        string       `gorm:"uniqueIndex" json:"username"`
	Email           string       `gorm:"uniqueIndex" json:"email"`
	Phone           string       `json:"phone"`
	Password        string       `json:"-"` // hide in response
	ProfileImage    ProfileImage `gorm:"embedded" json:"profileImage"`
	Role            string       `json:"role"` // user, admin, vendor, creator
	Bio             string       `json:"bio"`
	Country         string       `json:"country"`
	City            string       `json:"city"`
	Language        string       `json:"language"`
	Timezone        string       `json:"timezone"`
	IsVerified      bool         `json:"isVerified"`
	EmailVerifiedAt time.Time    `json:"emailVerifiedAt"`
	SocialLinks     SocialLinks  `gorm:"embedded" json:"socialLinks"`
	CreatedAt       time.Time    `json:"createdAt"`
	UpdatedAt       time.Time    `json:"updatedAt"`
}
