package main

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect DB
	database.Connect()

	log.Println("[INFO] Starting Database Seed...")

	// 1. Create Admin User
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)
	user := models.User{
		Name:     "Tri Wantoro",
		Username: "triwantoro",
		Email:    "ajarsinau@gmail.com",
		Password: string(hashedPassword),
		Role:     "admin",
		IsAdmin:  true,
	}
	database.DB.FirstOrCreate(&user, models.User{Email: "ajarsinau@gmail.com"})
	log.Printf("[OK] User created: %s", user.Email)

	// 2. Create Notebooks
	notebooks := []models.Notebook{
		{UserID: user.ID, Name: "Project TaraNote", Slug: "project-taranote"},
		{UserID: user.ID, Name: "Design Journey", Slug: "design-journey"},
		{UserID: user.ID, Name: "Changelog Story", Slug: "changelog-story"},
		{UserID: user.ID, Name: "RELEASE", Slug: "release"},
	}
	for i := range notebooks {
		database.DB.FirstOrCreate(&notebooks[i], models.Notebook{Slug: notebooks[i].Slug})
	}
	log.Println("[OK] Notebooks created")

	// 3. Create Sample Notes
	now := time.Now()
	notes := []models.Note{
		{
			UserID:      user.ID,
			NotebookID:  &notebooks[0].ID,
			Title:       "Welcome to TaraNote: A Digital Sanctuary",
			Slug:        "welcome-to-taranote-a-digital-sanctuary",
			Content:     "<h1>Welcome to TaraNote</h1><p>TaraNote is more than just a note-taking app; it is a digital sanctuary for your thoughts.</p><h3>Key Features:</h3><ul><li><strong>Zen Mode</strong>: A distraction-free writing environment.</li><li><strong>Seamless Publishing</strong>: Turn private notes into public articles with one click.</li><li><strong>Markdown Support</strong>: Write comfortably with familiar formatting.</li></ul>",
			Status:      "PUBLISHED",
			PublishedAt: &now,
		},
		{
			UserID:      user.ID,
			NotebookID:  &notebooks[0].ID,
			Title:       "The Philosophy of TaraNote",
			Slug:        "the-philosophy-of-taranote",
			Content:     "<h1>Why We Built This</h1><p>Many tools try to do everything. TaraNote does one thing well: <strong>Writing</strong>.</p><blockquote>\"The goal of design is not to make something beautiful. It's to make something beautiful that disappears.\"</blockquote>",
			Status:      "PUBLISHED",
			PublishedAt: &now,
		},
		{
			UserID:      user.ID,
			NotebookID:  &notebooks[1].ID,
			Title:       "The Zen Glass System (V6.5)",
			Slug:        "the-zen-glass-system-v6-5",
			Content:     "<h1>Enter Zen Mode</h1><p>Zen Mode wasn't a rejection of beauty. It was a refinement of <em>what kind</em> of beauty serves focus.</p><h3>The Philosophy Shift</h3><ul><li><strong>Electric Earth</strong> said: \"Look at me\"</li><li><strong>Zen Glass</strong> says: \"I'm here when you need me\"</li></ul>",
			Status:      "PUBLISHED",
			PublishedAt: &now,
		},
		{
			UserID:      user.ID,
			NotebookID:  &notebooks[3].ID,
			Title:       "Release v1.0.0",
			Slug:        "release-v1-0-0",
			Content:     "<h1>Production Release</h1><p>TaraNote v1.0.0 is now production ready. This release marks the transition from beta prototype to a stable platform.</p><h3>Architecture</h3><p>The backend has been rewritten in Go for blazing performance.</p>",
			Status:      "PUBLISHED",
			PublishedAt: &now,
		},
	}

	for i := range notes {
		var existingNote models.Note
		result := database.DB.Where("slug = ?", notes[i].Slug).First(&existingNote)
		if result.Error == nil {
			// Found, update it
			database.DB.Model(&existingNote).Updates(map[string]interface{}{
				"notebook_id":  notes[i].NotebookID,
				"title":        notes[i].Title,
				"content":      notes[i].Content,
				"status":       notes[i].Status,
				"published_at": notes[i].PublishedAt,
			})
			log.Printf("[OK] Note updated: %s", notes[i].Title)
		} else {
			// Not found, create it
			database.DB.Create(&notes[i])
			log.Printf("[OK] Note created: %s", notes[i].Title)
		}
	}
	log.Println("[OK] Notes seeded")

	log.Println("[DONE] Database Seed Completed!")
}
