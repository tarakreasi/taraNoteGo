package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"index" json:"user_id"`
	NotebookID  *uint          `gorm:"index;default:null" json:"notebook_id"`
	Title       string         `json:"title"`
	Slug        string         `gorm:"uniqueIndex" json:"slug"`
	Excerpt     string         `json:"excerpt"`
	Content     string         `json:"content"` // longText in GORM usually maps to string
	CoverImage  string         `json:"cover_image"`
	Status      string         `gorm:"default:'DRAFT'" json:"status"`
	PublishedAt *time.Time     `json:"published_at"`
	Views       int            `gorm:"default:0" json:"views"`
	IsFeatured  bool           `gorm:"default:false" json:"is_featured"` // Added based on seen migrations list earlier
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relations
	User     User      `json:"user,omitempty"`
	Notebook *Notebook `json:"notebook,omitempty"`
}
