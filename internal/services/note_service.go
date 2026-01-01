package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
)

type NoteService struct {
	validate *validator.Validate
}

func NewNoteService() *NoteService {
	return &NoteService{
		validate: validator.New(),
	}
}

type CreateNoteRequest struct {
	UserID     uint   `validate:"required"`
	Title      string `validate:"required,min=1,max=255"`
	NotebookID *uint
}

func (s *NoteService) CreateNote(req CreateNoteRequest) (*models.Note, error) {
	// 1. Validate
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// 2. Logic (Slug generation)
	slug := fmt.Sprintf("%s-%d", uuid.New().String(), time.Now().Unix())

	note := models.Note{
		UserID:     req.UserID,
		Title:      req.Title,
		Slug:       slug,
		Status:     "DRAFT",
		NotebookID: req.NotebookID,
	}

	// 3. Persistence
	if err := database.DB.Create(&note).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

type UpdateNoteRequest struct {
	ID         string `validate:"required"` // UUID or string ID
	UserID     uint   `validate:"required"`
	Title      string `validate:"omitempty,min=1"`
	Content    string
	Excerpt    string
	NotebookID *uint
	Status     string `validate:"omitempty,oneof=DRAFT PUBLISHED ARCHIVED"`
	IsFeatured bool
}

func (s *NoteService) UpdateNote(req UpdateNoteRequest) (*models.Note, error) {
	// 1. Validate
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// 2. Fetch Existing
	var note models.Note
	if err := database.DB.Where("id = ? AND user_id = ?", req.ID, req.UserID).First(&note).Error; err != nil {
		return nil, errors.New("note not found or unauthorized")
	}

	// 3. Update Fields
	note.Title = req.Title
	note.Content = req.Content
	note.Excerpt = req.Excerpt
	note.NotebookID = req.NotebookID
	note.Status = req.Status
	note.IsFeatured = req.IsFeatured

	// 4. Save
	if err := database.DB.Save(&note).Error; err != nil {
		return nil, err
	}

	return &note, nil
}

type NoteFilter struct {
	Search     string
	NotebookID *uint
	Status     string
}

func (s *NoteService) ListNotes(userID uint, filter NoteFilter) ([]models.Note, error) {
	var notes []models.Note
	db := database.DB.Preload("Notebook").Where("user_id = ?", userID)

	if filter.Search != "" {
		searchTerm := "%" + filter.Search + "%"
		db = db.Where("title LIKE ? OR content LIKE ?", searchTerm, searchTerm)
	}

	if filter.NotebookID != nil {
		db = db.Where("notebook_id = ?", *filter.NotebookID)
	}

	if filter.Status != "" {
		db = db.Where("status = ?", filter.Status)
	}

	// Order by updated_at desc by default
	db = db.Order("updated_at desc")

	if err := db.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (s *NoteService) DeleteNote(id string, userID uint) error {
	result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Note{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("note not found")
	}
	return nil
}
