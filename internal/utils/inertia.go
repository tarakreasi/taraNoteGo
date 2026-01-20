package utils

import (
	"encoding/json"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

// CreateInertiaPage generates the JSON object for the Inertia frontend
// CreateInertiaPage generates the JSON object for the Inertia frontend
func CreateInertiaPage(c *fiber.Ctx, component string, props fiber.Map) []byte {
	page := fiber.Map{
		"component": component,
		"props":     props,
		"url":       c.Path(),
		"version":   "v1", // Asset version, can be dynamic later
	}
	bytes, _ := json.Marshal(page)
	return bytes
}

// RenderInertia handles the Inertia response logic (JSON vs HTML)
func RenderInertia(c *fiber.Ctx, component string, props fiber.Map) error {
	// If X-Inertia header is present, return JSON
	if c.Get("X-Inertia") == "true" {
		c.Set("X-Inertia", "true")
		return c.JSON(fiber.Map{
			"component": component,
			"props":     props,
			"url":       c.Path(),
			"version":   "v1",
		})
	}

	// Otherwise render the full HTML page
	c.Set("Content-Type", "text/html; charset=utf-8")

	// Get Vite tags (Safe HTML)
	viteTags := GetViteTags()

	return c.Render("app", fiber.Map{
		"InertiaJSON": string(CreateInertiaPage(c, component, props)),
		"ViteTags":    template.HTML(viteTags),
	})
}
