package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tarakreasi/taraNote_go/internal/database"

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

	// Initialize View Engine
	engine := html.New("./views", ".html")

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "TaraNote Go v1.0",
		Views:   engine,
	})

	// Middleware
	app.Use(cors.New())

	// Static Assets (for built frontend)
	app.Static("/resources", "./resources")
	app.Static("/public", "./public")

	// Inertia Page Mock
	app.Get("/", func(c *fiber.Ctx) error {
		// This is a minimal definition of the Inertia Page object
		page := fiber.Map{
			"component": "TaraNote", // The Vue component name
			"props": fiber.Map{
				"auth": fiber.Map{
					"user": nil,
				},
				"can": fiber.Map{
					"login":    true,
					"register": true,
				},
			},
			"url":     "/",
			"version": "v1", // Asset version
		}

		pageBytes, _ := json.Marshal(page)

		return c.Render("app", fiber.Map{
			"InertiaJSON": string(pageBytes),
		})
	})

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
