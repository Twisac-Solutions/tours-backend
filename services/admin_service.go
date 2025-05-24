package services

import (
	"github.com/Twisac-Solutions/tours-backend/database"
	"github.com/Twisac-Solutions/tours-backend/models"
)

func FindAdminByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := database.DB.First(&admin, "email = ?", email).Error
	return &admin, err
}
