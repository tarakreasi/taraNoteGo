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
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// For API requests, return JSON
			if c.Get("Content-Type") == "application/json" || c.XHR() {
				return c.Status(code).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			// For View requests, render error page (or just text for now)
			// For View requests, render error page
			c.Set("Content-Type", "text/html")
			return c.Status(code).SendString("<!DOCTYPE html><html><head><title>Error</title></head><body><h1>Error</h1><p>" + err.Error() + "</p></body></html>")
		},
	})

	// Middleware
	app.Use(cors.New())
	// Custom Method Override Middleware
	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodPost {
			method := c.FormValue("_method")
			if method == "" {
				method = c.Get("X-HTTP-Method-Override")
			}
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
				c.Method(method)
			}
		}
		return c.Next()
	})

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
