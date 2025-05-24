package database

import (
	"fmt"

	"github.com/Twisac-Solutions/tours-backend/models"
	"github.com/Twisac-Solutions/tours-backend/utils"
)

func SeedSuperAdmin() {
	var count int64
	DB.Model(&models.User{}).Where("role = ?", "superadmin").Count(&count)
	if count > 0 {
		fmt.Println("✅ Superadmin already exists")
		return
	}

	superadmin := models.User{
		ID:       utils.GenerateUUID(),
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: utils.HashPassword("admin1234"),
		Role:     "superadmin",
	}

	if err := DB.Create(&superadmin).Error; err != nil {
		fmt.Println("❌ Failed to seed superadmin:", err)
		return
	}

	fmt.Println("✅ Superadmin seeded: email=admin@example.com, password=admin1234")
}
