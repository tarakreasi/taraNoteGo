package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect DB
	database.Connect()

	// Migrate
	log.Println("Running AutoMigrate...")
	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Notebook{},
		&models.Note{},
		&models.Setting{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}
