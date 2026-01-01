package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/utils"
)

// DocsView renders the documentation page
func DocsView(c *fiber.Ctx) error {
	path := c.Params("*")
	if path == "" {
		path = c.Params("path") // Fallback if using :path in some other context
	}
	if path == "" {
		path = "INDEX" // Default page
	}

	// 1. Sanitize Path (Prevent Directory Traversal)
	// Allow alphanumeric, underscores, hyphens, and slashes for subdirectories
	// But strictly prevent '..'
	if strings.Contains(path, "..") {
		return c.Status(400).SendString("Invalid path")
	}

	// 2. Construct absolute path to docs/
	// Assuming running from project root
	cwd, _ := os.Getwd()
	docPath := filepath.Join(cwd, "docs", path+".md")

	// 3. Check if file exists
	if _, err := os.Stat(docPath); os.IsNotExist(err) {
		// Try checking without extension or index
		docPathIndex := filepath.Join(cwd, "docs", path, "README.md")
		if _, err := os.Stat(docPathIndex); err == nil {
			docPath = docPathIndex
		} else {
			// Force HTML content type for 404 error page to avoid quirks mode
			c.Set("Content-Type", "text/html; charset=utf-8")
			c.Status(404)
			return utils.RenderInertia(c, "Docs", fiber.Map{
				"content":     "# 404 Not Found\n\nThe requested documentation page could not be found.",
				"currentPath": path,
				"displayName": "Not Found",
			})
		}
	}

	// 4. Read File Content
	contentBytes, err := os.ReadFile(docPath)
	if err != nil {
		return c.Status(500).SendString("Error reading documentation file")
	}

	content := string(contentBytes)

	// 5. Determine Display Name
	displayName := filepath.Base(path)
	displayName = strings.ReplaceAll(displayName, "_", " ")
	displayName = strings.Title(strings.ToLower(displayName))

	// 6. Get User (if logged in) for Floating Dock
	sess, _ := config.Store.Get(c)
	var user *models.User
	if sess != nil {
		userID := sess.Get("user_id")
		if userID != nil {
			var u models.User
			if err := database.DB.First(&u, userID).Error; err == nil {
				user = &u
			}
		}
	}

	props := fiber.Map{
		"content":     content,
		"currentPath": path,
		"displayName": displayName,
	}

	if user != nil {
		props["auth"] = fiber.Map{
			"user": user,
		}
	}

	return utils.RenderInertia(c, "Docs", props)
}
