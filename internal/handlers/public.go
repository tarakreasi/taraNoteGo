package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/utils"
)

// PublicList renders the home page with published articles
func PublicList(c *fiber.Ctx) error {
	var notes []models.Note

	// Fetch published notes, latest first
	// Preload User and Notebook for display
	if err := database.DB.Where("status = ?", "PUBLISHED").
		Preload("User").Preload("Notebook").
		Order("published_at desc").
		Limit(9). // Pagination limit
		Find(&notes).Error; err != nil {
		return c.Status(500).SendString("Error fetching notes")
	}

	// Mock settings for now (as seen in legacy controller)
	settings := fiber.Map{
		"articles_hero_title":    "Welcome to TaraNote",
		"articles_hero_subtitle": "Thoughts, stories, and ideas.",
		"footer_brand_name":      "TaraCreations",
		"footer_copyright_text":  "© 2025 TaraCreations. All rights reserved.",
	}

	return c.Render("app", fiber.Map{
		"InertiaJSON": string(utils.CreateInertiaPage(c, "Articles", fiber.Map{
			"notes": fiber.Map{
				"data": notes, // Inertia expects paginated 'data' usually
			},
			"topics":   []string{}, // Placeholder
			"settings": settings,
		})),
	})
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

	return c.Render("app", fiber.Map{
		"InertiaJSON": string(utils.CreateInertiaPage(c, "Article", fiber.Map{
			"article":  note,
			"settings": fiber.Map{},
		})),
	})
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

	// Mock notebooks
	var notebooks []models.Notebook
	database.DB.Find(&notebooks)

	return c.Render("app", fiber.Map{
		"InertiaJSON": string(utils.CreateInertiaPage(c, "TaraNote", fiber.Map{
			"notes":     notes,
			"notebooks": notebooks,
		})),
	})
}
