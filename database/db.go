package database

import (
	"log"
	"os"

	"github.com/Twisac-Solutions/tours-backend/models"
	libsql "github.com/ytsruh/gorm-libsql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var dialector gorm.Dialector

	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		dialector = libsql.New(libsql.Config{
			DSN:        dsn,
			DriverName: "libsql",
		})
	} else {
		// Local SQLite fallback
		dsn = os.Getenv("DB_PATH")
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
	DB.AutoMigrate(&models.User{},
		&models.Category{},
		&models.Destination{},
		&models.Tour{},
		&models.MediaDestination{},
		&models.MediaTour{})
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
