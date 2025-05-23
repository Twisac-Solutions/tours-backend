package database

import (
	"log"

	"github.com/your-username/tout-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("tout.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.User{})
}
