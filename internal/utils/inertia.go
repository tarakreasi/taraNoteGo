package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

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
