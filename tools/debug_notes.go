package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
)

func main() {
	godotenv.Load()
	database.Connect()

	var notes []models.Note
	if err := database.DB.Where("status = ?", "PUBLISHED").
		Limit(5).
		Find(&notes).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d published notes\n", len(notes))
	for _, n := range notes {
		fmt.Printf("ID: %d, Title: %s, Content Len: %d\n", n.ID, n.Title, len(n.Content))
		if len(n.Content) < 50 {
			fmt.Printf("Content: %s\n", n.Content)
		} else {
			fmt.Printf("Content (prefix): %s...\n", n.Content[:50])
		}
	}
}
