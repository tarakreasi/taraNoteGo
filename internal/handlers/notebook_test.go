package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/testutils"
	"golang.org/x/crypto/bcrypt"
)

func loginAndGetCookie(app *fiber.App, email, password string) string {
	// Helper to login and return session cookie
	resp, _, _ := testutils.MakeRequest(app, "POST", "/login", map[string]string{
		"email": email, "password": password,
	}, "")
	return resp.Header.Get("Set-Cookie")
}

func TestNotebook_CRUD(t *testing.T) {
	app := testutils.SetupApp()
	defer testutils.CleanupDB()

	// 1. Setup User
	password := "password"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Name: "Notebook User", Email: "nb@test.com", Password: string(hashed)}
	database.DB.Create(&user)

	// Login
	respLogin, _, err := testutils.MakeRequest(app, "POST", "/login", map[string]string{
		"email": "nb@test.com", "password": "password",
	}, "")
	assert.NoError(t, err)
	cookie := respLogin.Header.Get("Set-Cookie")
	assert.NotEmpty(t, cookie)

	// 2. Create Notebook
	payload := map[string]string{"name": "My First Notebook", "description": "Test Desc"}
	resp, body, err := testutils.MakeRequest(app, "POST", "/api/v1/admin/notebooks", payload, cookie)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createResult map[string]interface{}
	json.Unmarshal([]byte(body), &createResult)
	data := createResult["data"].(map[string]interface{})
	notebookID := uint(data["id"].(float64))

	assert.Equal(t, "My First Notebook", data["name"])
	assert.Equal(t, "my-first-notebook", data["slug"]) // Auto-generated slug

	// 3. List Notebooks
	respList, bodyList, err := testutils.MakeRequest(app, "GET", "/api/v1/admin/notebooks", nil, cookie)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, respList.StatusCode)

	var listResult map[string]interface{}
	json.Unmarshal([]byte(bodyList), &listResult)
	listData := listResult["data"].([]interface{})
	assert.Len(t, listData, 1)

	// 4. Update Notebook
	updatePayload := map[string]string{"name": "Updated Name"}
	respUpdate, bodyUpdate, err := testutils.MakeRequest(app, "PUT", fmt.Sprintf("/api/v1/admin/notebooks/%d", notebookID), updatePayload, cookie)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, respUpdate.StatusCode)

	var updateResult map[string]interface{}
	json.Unmarshal([]byte(bodyUpdate), &updateResult)
	updateData := updateResult["data"].(map[string]interface{})
	assert.Equal(t, "Updated Name", updateData["name"])
	assert.Equal(t, "updated-name", updateData["slug"]) // Auto-generated slug update

	// 5. Delete Notebook
	respDel, _, err := testutils.MakeRequest(app, "DELETE", fmt.Sprintf("/api/v1/admin/notebooks/%d", notebookID), nil, cookie)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, respDel.StatusCode)

	// Verify Deletion
	respList2, bodyList2, _ := testutils.MakeRequest(app, "GET", "/api/v1/admin/notebooks", nil, cookie)
	assert.Equal(t, http.StatusOK, respList2.StatusCode)
	json.Unmarshal([]byte(bodyList2), &listResult)
	listData2 := listResult["data"].([]interface{})
	assert.Len(t, listData2, 0)
}
