package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/handlers"
	"github.com/tarakreasi/taraNote_go/internal/middleware"
)

// SetupWeb routes
func SetupWeb(app *fiber.App) {
	// Guest Routes (Public)
	app.Get("/", handlers.PublicList)               // Home/Articles List
	app.Get("/articles/:slug", handlers.PublicShow) // Single Article
	app.Get("/taranote", handlers.TaraNoteBrowser)  // 3-Column Browser

	// Auth Routes
	app.Get("/login", handlers.ShowLogin)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	// Docs Routes
	app.Get("/docs", handlers.DocsView)
	app.Get("/docs/:path", handlers.DocsView)

	// Protected Routes
	app.Get("/dashboard", middleware.Protected, handlers.DashboardView)
}
