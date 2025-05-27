package config

import (
	"log"
	"go-voting-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("voting.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.Admin{},
		&models.Voter{},
		&models.Candidate{},
		&models.Election{},
		// &models.Report{},
		&models.Vote{}, // future use
	)

	if err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	DB = db
}
