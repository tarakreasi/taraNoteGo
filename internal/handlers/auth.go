package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/services"
	"github.com/tarakreasi/taraNote_go/internal/utils"
)

var authService = services.NewAuthService()

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

	// Use Service
	user, err := authService.Authenticate(services.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		// User not found or Invalid Password - Return 422 for Inertia
		c.Set("X-Inertia", "true")
		return c.Status(422).JSON(fiber.Map{
			"message": "Invalid credentials",
			"errors": fiber.Map{
				"email": "Invalid credentials",
			},
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
