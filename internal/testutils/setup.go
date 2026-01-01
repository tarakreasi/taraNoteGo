package testutils

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupApp initializes the app with an in-memory database for testing
func SetupApp() *fiber.App {
	// 1. Connect to In-Memory SQLite
	var err error
	database.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// 2. Migrate Schema
	database.DB.AutoMigrate(
		&models.User{},
		&models.Note{},
		&models.Notebook{},
		&models.Setting{},
	)

	// 3. Init Session (Store)
	config.InitSession()

	// 4. Setup Fiber
	// Note: We use the actual view engine for complete integration, 
	// but point to the real views directory relative to where tests run (usually from internal/handlers)
	// We might need to adjust path depending on where tests are run
	engine := html.New("../../views", ".html") 

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// 5. Setup Routes
	routes.SetupWeb(app)
	routes.SetupAPI(app)

	return app
}

// CleanupDB clears the database (optional if using :memory: and restarting app, but good for between tests)
func CleanupDB() {
	database.DB.Exec("DELETE FROM users")
	database.DB.Exec("DELETE FROM notes")
	database.DB.Exec("DELETE FROM notebooks")
	database.DB.Exec("DELETE FROM settings")
}
