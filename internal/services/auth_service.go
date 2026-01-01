package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	validate *validator.Validate
}

func NewAuthService() *AuthService {
	return &AuthService{
		validate: validator.New(),
	}
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (s *AuthService) Authenticate(req LoginRequest) (*models.User, error) {
	// 1. Validate Input
	if err := s.validate.Struct(req); err != nil {
		return nil, err
	}

	// 2. Find User
	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 3. Verify Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
