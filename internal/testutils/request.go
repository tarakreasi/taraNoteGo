package testutils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gofiber/fiber/v2"
)

// MakeRequest sends a request to the app and returns the response
// method: GET, POST, etc.
// url: /login, /api/notes
// body: nil or struct (will be JSON encoded)
// cookies: optional string (for session)
func MakeRequest(app *fiber.App, method, url string, body interface{}, cookies string) (*http.Response, string, error) {
	var reqBody io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, "", err
		}
		reqBody = bytes.NewBuffer(jsonBytes)
	}

	req := httptest.NewRequest(method, url, reqBody)
	req.Header.Set("Content-Type", "application/json")
	if cookies != "" {
		req.Header.Set("Cookie", cookies)
	}

	resp, err := app.Test(req, -1) // -1 disables timeout
	if err != nil {
		return nil, "", err
	}

	// Read response body
	respBodyBytes, _ := io.ReadAll(resp.Body)
	respString := string(respBodyBytes)

	return resp, respString, nil
}
