package database

import (
	"fmt"
	"log"
)

// MigrateDB runs all database migrations
func MigrateDB() {
	// Run specific migrations
	removeDescColumnFromTours()
}

// removeDescColumnFromTours checks if 'desc' column exists in tours table and removes it
func removeDescColumnFromTours() {
	// Check if the 'desc' column exists in the tours table
	var count int64
	err := DB.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = 'tours' AND COLUMN_NAME = 'desc'").Count(&count).Error
	if err != nil {
		log.Printf("Error checking if 'desc' column exists: %v", err)
		return
	}

	// If the column exists, drop it
	if count > 0 {
		// For PostgreSQL
		if err := DB.Exec("ALTER TABLE tours DROP COLUMN IF EXISTS \"desc\"").Error; err != nil {
			log.Printf("Error dropping 'desc' column from tours table: %v", err)
			return
		}

		// For SQLite
		if err := DB.Exec("PRAGMA foreign_keys = OFF").Error; err == nil {
			if err := DB.Exec("ALTER TABLE tours DROP COLUMN desc").Error; err != nil {
				log.Printf("Error dropping 'desc' column from tours table (SQLite): %v", err)
			}
			DB.Exec("PRAGMA foreign_keys = ON")
		}

		fmt.Println("✅ Removed 'desc' column from tours table")
	} else {
		fmt.Println("ℹ️ 'desc' column does not exist in tours table")
	}
}
