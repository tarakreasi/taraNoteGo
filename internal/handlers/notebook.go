package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
)

// ListNotebooks returns all notebooks for the authenticated user
func ListNotebooks(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")

	var notebooks []models.Notebook
	if err := database.DB.Where("user_id = ?", userID).Find(&notebooks).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch notebooks"})
	}

	return c.JSON(notebooks)
}

// CreateNotebook creates a new notebook
func CreateNotebook(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id").(uint)

	type CreateRequest struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	}

	req := new(CreateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	notebook := models.Notebook{
		UserID:      userID,
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	}

	if err := database.DB.Create(&notebook).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create notebook"})
	}

	return c.Status(201).JSON(notebook)
}

// UpdateNotebook updates an existing notebook
func UpdateNotebook(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")
	id := c.Params("id")

	var notebook models.Notebook
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&notebook).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Notebook not found"})
	}

	type UpdateRequest struct {
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
	}

	req := new(UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	notebook.Name = req.Name
	notebook.Slug = req.Slug
	notebook.Description = req.Description

	if err := database.DB.Save(&notebook).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update notebook"})
	}

	return c.JSON(notebook)
}

// DeleteNotebook deletes a notebook
func DeleteNotebook(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")
	id := c.Params("id")

	result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Notebook{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete notebook"})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Notebook not found"})
	}

	return c.SendStatus(204)
}
