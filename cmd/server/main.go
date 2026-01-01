package main

import (
	"log"
	"os"

	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/handlers"
	"github.com/tarakreasi/taraNote_go/internal/middleware"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/utils"

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

	// Guest Routes
	// Guest Routes (Public)
	app.Get("/", handlers.PublicList)               // Home/Articles List
	app.Get("/articles/:slug", handlers.PublicShow) // Single Article
	app.Get("/taranote", handlers.TaraNoteBrowser)  // 3-Column Browser

	// Auth Routes
	// Auth Routes
	app.Get("/login", handlers.ShowLogin)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	// Docs Routes
	app.Get("/docs", handlers.DocsView)
	app.Get("/docs/:path", handlers.DocsView)
	app.Post("/logout", handlers.Logout)

	// Protected Routes
	app.Get("/dashboard", middleware.Protected, func(c *fiber.Ctx) error {
		// Get User ID from session
		sess, _ := config.Store.Get(c)
		userID := sess.Get("user_id")

		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			return c.Redirect("/login")
		}

		return c.Render("app", fiber.Map{
			"InertiaJSON": string(utils.CreateInertiaPage(c, "Dashboard", fiber.Map{
				"auth": fiber.Map{
					"user": user,
				},
			})),
		})
	})

	// API Group (Protected)
	api := app.Group("/api/v1/admin", middleware.Protected)
	api.Get("/notebooks", handlers.ListNotebooks)
	api.Post("/notebooks", handlers.CreateNotebook)
	api.Put("/notebooks/:id", handlers.UpdateNotebook)
	api.Delete("/notebooks/:id", handlers.DeleteNotebook)

	// Notes
	api.Get("/notes", handlers.ListNotes)
	api.Post("/notes", handlers.CreateNote)
	api.Put("/notes/:id", handlers.UpdateNote)
	api.Delete("/notes/:id", handlers.DeleteNote)

	// Uploads
	api.Post("/upload", handlers.UploadImage)

	// Settings
	api.Get("/settings", handlers.ListSettings)
	api.Post("/settings", handlers.UpdateSettings)

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
