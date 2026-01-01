package handlers

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tarakreasi/taraNote_go/internal/config"
	"github.com/tarakreasi/taraNote_go/internal/services"
)

var noteService = services.NewNoteService()

// ListNotes returns all notes
func ListNotes(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id")

	notes, err := noteService.ListNotes(userID.(uint))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch notes"})
	}

	return c.JSON(fiber.Map{"data": notes})
}

// CreateNote creates a new draft note
func CreateNote(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id").(uint)

	type Request struct {
		Title string `json:"title"`
	}

	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	note, err := noteService.CreateNote(services.CreateNoteRequest{
		UserID: userID,
		Title:  req.Title,
	})

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(fiber.Map{"data": note})
}

// UpdateNote updates a note content
func UpdateNote(c *fiber.Ctx) error {
	sess, _ := config.Store.Get(c)
	userID := sess.Get("user_id").(uint)
	id := c.Params("id")

	type Request struct {
		Title      string `json:"title"`
		Content    string `json:"content"`
		Excerpt    string `json:"excerpt"`
		NotebookID *uint  `json:"notebook_id"`
		Status     string `json:"status"`
		IsFeatured bool   `json:"is_featured"`
	}

	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	note, err := noteService.UpdateNote(services.UpdateNoteRequest{
		ID:         id,
		UserID:     userID,
		Title:      req.Title,
		Content:    req.Content,
		Excerpt:    req.Excerpt,
		NotebookID: req.NotebookID,
		Status:     req.Status,
		IsFeatured: req.IsFeatured,
	})

	if err != nil {
		status := 500
		if err.Error() == "note not found or unauthorized" {
			status = 404
		} else if err != nil {
			// Validation errors, etc
			status = 400
		}
		return c.Status(status).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"data": note})
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
	userID := sess.Get("user_id").(uint)
	id := c.Params("id")

	err := noteService.DeleteNote(id, userID)
	if err != nil {
		if err.Error() == "note not found" {
			return c.Status(404).JSON(fiber.Map{"error": "Note not found"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete note"})
	}

	return c.SendStatus(204)
}
