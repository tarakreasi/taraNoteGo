package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/handlers"
	"github.com/tarakreasi/taraNote_go/internal/middleware"
)

// SetupWeb routes
func SetupWeb(app *fiber.App) {
	// Guest Routes (Public)
	app.Get("/", handlers.PublicList).Name("home")                        // Home/Articles List
	app.Get("/articles/:slug", handlers.PublicShow).Name("articles.show") // Single Article
	app.Get("/taranote", handlers.TaraNoteBrowser).Name("taranote")       // 3-Column Browser

	// Auth Routes
	app.Get("/login", handlers.ShowLogin).Name("login.view")
	app.Post("/login", handlers.Login).Name("login.post")
	app.Post("/logout", handlers.Logout).Name("logout")

	// Docs Routes
	app.Get("/docs", handlers.DocsView).Name("docs.index")
	app.Get("/docs/:path", handlers.DocsView).Name("docs.show")

	// Protected Routes
	app.Get("/dashboard", middleware.Protected, handlers.DashboardView).Name("dashboard")
}
