package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Manifest structure for Vite build manifest
type Manifest map[string]struct {
	File    string   `json:"file"`
	Src     string   `json:"src"`
	IsEntry bool     `json:"isEntry"`
	Css     []string `json:"css"`
}

// GetViteTags generates the HTML tags for Vite assets
// In development, it returns the Vite dev server client and app entry.
// In production, it reads the manifest.json and returns the built assets.
func GetViteTags() string {
	appEnv := os.Getenv("APP_ENV")
	if appEnv != "production" {
		return `
		<script type="module" src="http://localhost:5173/@vite/client"></script>
        <script type="module" src="http://localhost:5173/resources/js/app.js"></script>
		`
	}

	// Production: Read manifest
	manifestPath := "public/build/manifest.json"
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Sprintf("<!-- Error reading manifest: %v -->", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(data, &manifest); err != nil {
		return fmt.Sprintf("<!-- Error parsing manifest: %v -->", err)
	}

	// Find the entry point (assuming resources/js/app.js)
	entryKey := "resources/js/app.js"
	entry, ok := manifest[entryKey]
	if !ok {
		return "<!-- Entry point not found in manifest -->"
	}

	tags := fmt.Sprintf(`<script type="module" src="/build/%s"></script>`, entry.File)

	// Add CSS
	for _, css := range entry.Css {
		tags += fmt.Sprintf(`<link rel="stylesheet" href="/build/%s">`, css)
	}

	return tags
}
