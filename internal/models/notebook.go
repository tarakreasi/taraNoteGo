package models

import (
	"time"

	"gorm.io/gorm"
)

type Notebook struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"index" json:"user_id"`
	Name        string         `json:"name"`
	Slug        string         `gorm:"uniqueIndex" json:"slug"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Relations
	User User `json:"user,omitempty"`
}
