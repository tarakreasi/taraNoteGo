package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"github.com/tarakreasi/taraNote_go/internal/testutils"
	"golang.org/x/crypto/bcrypt"
)

func TestNote_CRUD(t *testing.T) {
	app := testutils.SetupApp()
	defer testutils.CleanupDB()

	// 1. Setup User & Notebook
	password := "password"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{Name: "Note User", Email: "note@test.com", Password: string(hashed)}
	database.DB.Create(&user)

	notebook := models.Notebook{UserID: user.ID, Name: "Test Notebook", Slug: "test-notebook"}
	database.DB.Create(&notebook)

	// Login
	respLogin, _, err := testutils.MakeRequest(app, "POST", "/login", map[string]string{
		"email": "note@test.com", "password": "password",
	}, "")
	assert.NoError(t, err)
	cookie := respLogin.Header.Get("Set-Cookie")

	// 2. Create Note
	payload := map[string]interface{}{
		"title":       "My New Note",
		"content":     "<p>Hello</p>",
		"status":      "DRAFT",
		"notebook_id": notebook.ID,
	}
	resp, body, err := testutils.MakeRequest(app, "POST", "/api/v1/admin/notes", payload, cookie)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	var createResult map[string]interface{}
	json.Unmarshal([]byte(body), &createResult)
	data := createResult["data"].(map[string]interface{})
	noteID := uint(data["id"].(float64))

	assert.Equal(t, "My New Note", data["title"])
	assert.NotEmpty(t, data["slug"])

	// 3. List Notes (Filter)
	// Filter by Status DRAFT
	respList, bodyList, _ := testutils.MakeRequest(app, "GET", "/api/v1/admin/notes?status=DRAFT", nil, cookie)
	assert.Equal(t, http.StatusOK, respList.StatusCode)
	var listResult map[string]interface{}
	json.Unmarshal([]byte(bodyList), &listResult)
	listData := listResult["data"].([]interface{})
	assert.Len(t, listData, 1)

	// Filter by Status PUBLISHED (Should be empty)
	respList2, bodyList2, _ := testutils.MakeRequest(app, "GET", "/api/v1/admin/notes?status=PUBLISHED", nil, cookie)
	assert.Equal(t, http.StatusOK, respList2.StatusCode)
	json.Unmarshal([]byte(bodyList2), &listResult)
	listData2 := listResult["data"].([]interface{})
	assert.Len(t, listData2, 0)

	// 4. Update Note (PUT)
	updatePayload := map[string]interface{}{
		"title":       "Updated Note Title",
		"status":      "PUBLISHED",
		"notebook_id": notebook.ID,
	}
	respUpdate, _, err := testutils.MakeRequest(app, "PUT", fmt.Sprintf("/api/v1/admin/notes/%d", noteID), updatePayload, cookie)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, respUpdate.StatusCode)

	// Verify Update
	var note models.Note
	database.DB.First(&note, noteID)
	assert.Equal(t, "Updated Note Title", note.Title)
	assert.Equal(t, "PUBLISHED", note.Status)

	// 5. User Isolation Check
	// Create another user
	hashedP, _ := bcrypt.GenerateFromPassword([]byte("p"), bcrypt.DefaultCost)
	otherUser := models.User{Name: "Other", Email: "other@test.com", Password: string(hashedP)}
	database.DB.Create(&otherUser)
	// Login as other user
	respLogin2, _, _ := testutils.MakeRequest(app, "POST", "/login", map[string]string{
		"email": "other@test.com", "password": "p",
	}, "")
	cookie2 := respLogin2.Header.Get("Set-Cookie")

	// Try to list notes (Should see 0)
	respList3, bodyList3, _ := testutils.MakeRequest(app, "GET", "/api/v1/admin/notes", nil, cookie2)
	assert.Equal(t, http.StatusOK, respList3.StatusCode)
	json.Unmarshal([]byte(bodyList3), &listResult)
	listData3 := listResult["data"].([]interface{})
	assert.Len(t, listData3, 0)
}
