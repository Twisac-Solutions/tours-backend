package database

import (
	"fmt"
	"log"
)

// MigrateDB runs all database migrations
func MigrateDB() {
	// Run specific migrations
	removeDescColumnFromTours()
	addRatingFieldsToTours()
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

// addRatingFieldsToTours adds rating fields to the tours table
func addRatingFieldsToTours() {
	// Check if the average_rating column exists
	var count int64
	err := DB.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = 'tours' AND COLUMN_NAME = 'average_rating'").Count(&count).Error
	if err != nil {
		log.Printf("Error checking if 'average_rating' column exists: %v", err)
		return
	}

	// If the column doesn't exist, add it
	if count == 0 {
		// For PostgreSQL
		if err := DB.Exec("ALTER TABLE tours ADD COLUMN average_rating DECIMAL(3,2) DEFAULT 0.00").Error; err != nil {
			log.Printf("Error adding 'average_rating' column to tours table: %v", err)
		}

		// For SQLite
		if err := DB.Exec("ALTER TABLE tours ADD COLUMN average_rating REAL DEFAULT 0.0").Error; err != nil {
			log.Printf("Error adding 'average_rating' column to tours table (SQLite): %v", err)
		}
	}

	// Check if the review_count column exists
	count = 0
	err = DB.Raw("SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = 'tours' AND COLUMN_NAME = 'review_count'").Count(&count).Error
	if err != nil {
		log.Printf("Error checking if 'review_count' column exists: %v", err)
		return
	}

	// If the column doesn't exist, add it
	if count == 0 {
		// For PostgreSQL
		if err := DB.Exec("ALTER TABLE tours ADD COLUMN review_count INTEGER DEFAULT 0").Error; err != nil {
			log.Printf("Error adding 'review_count' column to tours table: %v", err)
		}

		// For SQLite
		if err := DB.Exec("ALTER TABLE tours ADD COLUMN review_count INTEGER DEFAULT 0").Error; err != nil {
			log.Printf("Error adding 'review_count' column to tours table (SQLite): %v", err)
		}
	}

	fmt.Println("✅ Added rating fields to tours table")
}
