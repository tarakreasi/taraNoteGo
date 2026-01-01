package handlers_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"

	"github.com/tarakreasi/taraNote_go/internal/testutils"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin_Success(t *testing.T) {
	app := testutils.SetupApp()
	defer testutils.CleanupDB()

	// 1. Seed User
	password := "password123"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: string(hashed),
	}
	database.DB.Create(&user)

	// 2. Make Request
	loginPayload := map[string]string{
		"email":    "test@example.com",
		"password": password,
	}
	resp, _, err := testutils.MakeRequest(app, "POST", "/login", loginPayload, "")

	// 3. Assertions
	assert.NoError(t, err)
	// Expect Redirect to Dashboard (302)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, "/dashboard", resp.Header.Get("Location"))
}

func TestLogin_InvalidCredentials(t *testing.T) {
	app := testutils.SetupApp()
	defer testutils.CleanupDB()

	loginPayload := map[string]string{
		"email":    "wrong@example.com",
		"password": "wrongpassword",
	}
	resp, _, err := testutils.MakeRequest(app, "POST", "/login", loginPayload, "")

	assert.NoError(t, err)
	// Expect 401 or 422
	assert.NotEqual(t, http.StatusOK, resp.StatusCode)
}
