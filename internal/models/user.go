package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Name             string         `json:"name"`
	Username         string         `json:"username"`
	Email            string         `gorm:"uniqueIndex" json:"email"`
	EmailVerifiedAt  *time.Time     `json:"email_verified_at"`
	Password         string         `json:"-"` // Don't expose password in JSON
	RememberToken    string         `json:"remember_token"`
	ProfilePhotoPath string         `json:"profile_photo_path"`
	Role             string         `json:"role"` // 'admin', 'editor', 'user'
	IsAdmin          bool           `json:"is_admin"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// CheckPassword compares the provided password with the stored hash
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// HashPassword hashes the password using bcrypt
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}
