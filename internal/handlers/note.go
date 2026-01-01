package handlers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
)

// ListNotes returns all notes (with search/pagination later)
func ListNotes(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")

	var notes []models.Note
	// GORM preload notebook
	if err := database.DB.Preload("Notebook").Where("user_id = ?", userID).Find(&notes).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch notes"})
	}

	return c.JSON(notes)
}

// CreateNote creates a new draft note
func CreateNote(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id").(uint)

	type CreateRequest struct {
		Title string `json:"title"`
	}

	req := new(CreateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	// Generate basic slug
	slug := fmt.Sprintf("%s-%d", uuid.New().String(), time.Now().Unix())

	note := models.Note{
		UserID: userID,
		Title:  req.Title, // Or "Untitled"
		Slug:   slug,
		Status: "DRAFT",
	}

	if err := database.DB.Create(&note).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create note"})
	}

	return c.Status(201).JSON(note)
}

// UpdateNote updates a note content
func UpdateNote(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")
	id := c.Params("id")

	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
	}

	type UpdateRequest struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		Excerpt    string `json:"excerpt"`
		NotebookID *uint  `json:"notebook_id"`
		Status     string `json:"status"`
		IsFeatured bool   `json:"is_featured"`
	}

	req := new(UpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	note.Title = req.Title
	note.Content = req.Content
	note.Excerpt = req.Excerpt
	note.NotebookID = req.NotebookID
	note.Status = req.Status
	note.IsFeatured = req.IsFeatured

	if err := database.DB.Save(&note).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update note"})
	}

	return c.JSON(note)
}

// UploadImage handles image uploads from Tiptap or Cover Image
func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}

	// Generate unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)

	// Save to public uploads
	path := fmt.Sprintf("./resources/public/uploads/%s", filename)

	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}

	// Return URL
	url := fmt.Sprintf("/public/uploads/%s", filename)
	return c.JSON(fiber.Map{"url": url})
}

func DeleteNote(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")
	id := c.Params("id")

	result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{})
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete note"})
	}

	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
	}

	return c.SendStatus(204)
}
