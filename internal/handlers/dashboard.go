package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/utils"
)

// DashboardView renders the dashboard page
func DashboardView(c *fiber.Ctx) error {
	// Get User ID from session
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return c.Redirect("/login")
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Redirect("/login")
	}

	return c.Render("app", fiber.Map{
		"InertiaJSON": string(utils.CreateInertiaPage(c, "Dashboard", fiber.Map{
			"auth": fiber.Map{
				"user": user,
			},
		})),
	})
}
