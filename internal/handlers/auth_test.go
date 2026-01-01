package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/testutils"
	"golang.org/x/crypto/bcrypt"
)

func TestAuth_Login_Success(t *testing.T) {
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
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, "/dashboard", resp.Header.Get("Location"))

	// 4. Verify Cookie
	cookies := resp.Cookies()
	assert.NotEmpty(t, cookies)
	assert.Equal(t, "session_id", cookies[0].Name)
}

func TestAuth_Login_InvalidCredentials(t *testing.T) {
	app := testutils.SetupApp()
	defer testutils.CleanupDB()

	loginPayload := map[string]string{
		"email":    "wrong@example.com",
		"password": "wrongpassword",
	}
	resp, body, err := testutils.MakeRequest(app, "POST", "/login", loginPayload, "")

	assert.NoError(t, err)
	// Expect 422 Unprocessable Entity (Inertia Validation Error)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)

	// Check Inertia Header
	assert.Equal(t, "true", resp.Header.Get("X-Inertia"))

	// Check JSON Body
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	assert.Contains(t, result["errors"], "email")
}

func TestAuth_Logout(t *testing.T) {
	app := testutils.SetupApp()

	// Simulate logged in state (requires session mocking or bypass)
	// For integration test, we login first then logout
	password := "pass"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Email: "logout@example.com", Password: string(hashed)}
	database.DB.Create(&user)

	// Login to get cookie
	respLogin, _, _ := testutils.MakeRequest(app, "POST", "/login", map[string]string{
		"email": "logout@example.com", "password": "pass",
	}, "")
	cookie := respLogin.Header.Get("Set-Cookie")

	// Logout
	resp, _, err := testutils.MakeRequest(app, "POST", "/logout", nil, cookie)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusFound, resp.StatusCode)
	assert.Equal(t, "/", resp.Header.Get("Location"))
}
