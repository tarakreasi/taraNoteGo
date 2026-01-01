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
	api.Get("/notebooks", handlers.ListNotebooks).Name("api.notebooks.index")
	api.Post("/notebooks", handlers.CreateNotebook).Name("api.notebooks.store")
	api.Put("/notebooks/:id", handlers.UpdateNotebook).Name("api.notebooks.update")
	api.Delete("/notebooks/:id", handlers.DeleteNotebook).Name("api.notebooks.destroy")

	// Notes
	api.Get("/notes", handlers.ListNotes).Name("api.notes.index")
	api.Post("/notes", handlers.CreateNote).Name("api.notes.store")
	api.Put("/notes/:id", handlers.UpdateNote).Name("api.notes.update")
	api.Delete("/notes/:id", handlers.DeleteNote).Name("api.notes.destroy")

	// Uploads
	api.Post("/upload", handlers.UploadImage).Name("api.upload")

	// Settings
	api.Get("/settings", handlers.ListSettings).Name("api.settings.index")
	api.Post("/settings", handlers.UpdateSettings).Name("api.settings.update")
}
