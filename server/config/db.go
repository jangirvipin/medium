package config

import (
	"fmt"
	"log"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"medium/server/models"
)

// ConnectDatabase initializes and returns the database connection
func ConnectDatabase() *gorm.DB {
	fmt.Println("Connecting to database...")
	dsn:=os.Getenv("URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate models
	err = db.AutoMigrate(&models.User{},&models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate models:", err)
	}

	fmt.Println("Database connection established")
	return db
}
