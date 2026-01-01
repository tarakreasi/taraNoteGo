package main

import (
	"log"
	"os"

	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to Database
	database.Connect()

	// Initialize Session
	config.InitSession()

	// Initialize View Engine
	engine := html.New("./views", ".html")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "TaraNote Go v1.0",
		Views:   engine,
	})

	// Middleware
	app.Use(cors.New())

	// Static Assets
	app.Static("/resources", "./resources")
	app.Static("/public", "./public")
	app.Static("/images", "./public/images")

	// --- ROUTES ---
	routes.SetupWeb(app)
	routes.SetupAPI(app)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
