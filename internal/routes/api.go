package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/handlers"
	"github.com/tarakreasi/taraNote_go/internal/middleware"
)

// SetupAPI routes
func SetupAPI(app *fiber.App) {
	// API Group (Protected)
	api := app.Group("/api/v1/admin", middleware.Protected)

	// Notebooks
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
}
