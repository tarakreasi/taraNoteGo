package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/utils"
)

// Helper to get authenticated user
func getAuthUser(c *fiber.Ctx) *models.User {
	sess, _ := config.Store.Get(c)
	if sess != nil {
		userID := sess.Get("user_id")
		if userID != nil {
			var user models.User
			if err := database.DB.First(&user, userID).Error; err == nil {
				return &user
			}
		}
	}
	return nil
}

// PublicList renders the home page with published articles
func PublicList(c *fiber.Ctx) error {
	var notes []models.Note

	// Fetch published notes, latest first
	if err := database.DB.Where("status = ?", "PUBLISHED").
		Preload("User").Preload("Notebook").
		Order("published_at desc").
		Limit(9).
		Find(&notes).Error; err != nil {
		return c.Status(500).SendString("Error fetching notes")
	}

	// Fetch notebooks for sidebar
	var notebooks []models.Notebook
	database.DB.Find(&notebooks)

	props := fiber.Map{
		"notes":     notes,
		"notebooks": notebooks,
	}

	// Add Auth User
	if user := getAuthUser(c); user != nil {
		props["auth"] = fiber.Map{"user": user}
	}

	return utils.RenderInertia(c, "TaraNote", props)
}

// PublicShow renders a single article by slug
func PublicShow(c *fiber.Ctx) error {
	slug := c.Params("slug")
	var note models.Note

	// Find note by slug
	if err := database.DB.Where("slug = ?", slug).
		Preload("User").Preload("Notebook").
		First(&note).Error; err != nil {
		// If not found, return 404 Inertia page or simple 404
		return c.Status(404).SendString("Article not found")
	}

	// Check if published (unless owner - logic simplified for public view)
	if note.Status != "PUBLISHED" {
		return c.Status(404).SendString("Article not found")
	}

	props := fiber.Map{
		"article":  note,
		"settings": fiber.Map{},
	}

	// Add Auth User
	if user := getAuthUser(c); user != nil {
		props["auth"] = fiber.Map{"user": user}
	}

	return utils.RenderInertia(c, "Docs", props)
}

// TaraNoteBrowser renders the 3-column internal browser
func TaraNoteBrowser(c *fiber.Ctx) error {
	var notes []models.Note
	if err := database.DB.Where("status = ?", "PUBLISHED").
		Preload("User").Preload("Notebook").
		Order("published_at desc").
		Find(&notes).Error; err != nil {
		return c.Status(500).SendString("Error fetching notes")
	}

	// Fetch notebooks with note counts
	var notebooks []models.Notebook
	database.DB.Select("notebooks.*, count(notes.id) as notes_count").
		Joins("LEFT JOIN notes ON notes.notebook_id = notebooks.id AND notes.deleted_at IS NULL").
		Group("notebooks.id").
		Find(&notebooks)

	props := fiber.Map{
		"notes":     notes,
		"notebooks": notebooks,
	}

	// Add Auth User
	if user := getAuthUser(c); user != nil {
		props["auth"] = fiber.Map{"user": user}
	}

	return utils.RenderInertia(c, "TaraNote", props)
}
