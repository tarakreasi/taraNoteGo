package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
)

// Protected ensures the user is logged in
func Protected(c *fiber.Ctx) error {
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Status(500).SendString("Session error")
	}

	if sess.Get("user_id") == nil {
		// If API request, return 401
		if c.Path()[:4] == "/api" || c.Get("Content-Type") == "application/json" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}
		// Else, redirect to login
		return c.Redirect("/login")
	}

	return c.Next()
}
