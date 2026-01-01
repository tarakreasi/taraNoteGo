package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/utils"
)

// ShowLogin renders the login page (Inertia)
func ShowLogin(c *fiber.Ctx) error {
	return c.Render("app", fiber.Map{
		"InertiaJSON": string(utils.CreateInertiaPage(c, "Auth/Login", fiber.Map{})),
	})
}

// Login handles the authentication attempt
func Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).SendString("Bad Request")
	}

	var user models.User
	// Find user by email
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		// User not found - Return 422 for Inertia
		return c.Status(422).JSON(fiber.Map{
			"message": "Invalid credentials",
			"errors": fiber.Map{
				"email": "Invalid credentials",
			},
		})
	}

	// Check password
	if !user.CheckPassword(req.Password) {
		return c.Status(422).JSON(fiber.Map{
			"message": "Invalid credentials",
			"errors": fiber.Map{
				"email": "Invalid credentials"},
		})
	}

	// Create Session
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	sess.Set("user_id", user.ID)
	if err := sess.Save(); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Redirect to Dashboard
	return c.Redirect("/dashboard")
}

// Logout destroys the session
func Logout(c *fiber.Ctx) error {
	sess, err := config.Store.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if err := sess.Destroy(); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Redirect("/")
}
