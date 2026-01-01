package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	// Default to SQLite for now, mirroring the Laravel setup if it uses sqlite
	// In production this might be different.
	dbName := os.Getenv("DB_DATABASE")
	if dbName == "" {
		dbName = "database/database.sqlite"
	}

	// Check if we are using sqlite
	// For now we assume sqlite for simplicity of migration startup
	// You can add Postgres/MySQL drivers here as needed.

	// Enable WAL mode and busy timeout for concurrent writes
	dsn := fmt.Sprintf("%s?_journal_mode=WAL&_busy_timeout=5000&_synchronous=NORMAL", dbName)
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successfully opened")
}
