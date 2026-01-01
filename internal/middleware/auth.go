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
		return c.Redirect("/login")
	}

	return c.Next()
}
