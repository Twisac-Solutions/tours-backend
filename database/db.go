package database

import (
	"log"

	"github.com/Twisac-Solutions/tours-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tour.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Destination{})
	// DB.AutoMigrate(
	// 		&models.User{},
	// 		&models.Admin{},
	// 		&models.Category{},
	// 		&models.Destination{},
	// 		&models.Event{},
	// 		&models.Media{},
	// 		&models.Review{},
	// 		&models.Tag{},
	// 		&models.Tour{},
	// 	)
	if err != nil {
		log.Fatal("Failed to auto-migrate:", err)
	}
	log.Println("Auto-migration completed successfully")
}
