package database

import (
	"log"
	"os"

	"github.com/Twisac-Solutions/tours-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var dialector gorm.Dialector

	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		// production / any server
		dialector = postgres.Open(dsn)
	} else {
		// local dev (SQLite)
		dsn := os.Getenv("DB_PATH")
		if dsn == "" {
			dsn = "tour.db"
		}
		dialector = sqlite.Open(dsn)
	}

	var err error
	DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	if err := DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Destination{},
		&models.Tour{},
		&models.Review{},
		&models.MediaDestination{},
		&models.MediaTour{},
	); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}
}
