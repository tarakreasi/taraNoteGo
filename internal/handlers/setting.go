package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tarakreasi/taraNote_go/internal/database"
	"github.com/tarakreasi/taraNote_go/internal/models"
)

// ListSettings returns all settings as a key-value map or list
func ListSettings(c *fiber.Ctx) error {
	var settings []models.Setting
	if err := database.DB.Find(&settings).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch settings"})
	}
	return c.JSON(settings)
}

// UpdateSettings updates multiple settings at once
func UpdateSettings(c *fiber.Ctx) error {
	type SettingUpdate struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	var updates []SettingUpdate
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}

	// Transaction for bulk update
	tx := database.DB.Begin()

	for _, update := range updates {
		var setting models.Setting
		if err := tx.Where("key = ?", update.Key).First(&setting).Error; err != nil {
			// Create if not exists (upsert-ish)
			newSetting := models.Setting{
				Key:   update.Key,
				Value: update.Value,
				Type:  "text", // Default
			}
			if err := tx.Create(&newSetting).Error; err != nil {
				tx.Rollback()
				return c.Status(500).JSON(fiber.Map{"error": "Failed to create setting " + update.Key})
			}
		} else {
			setting.Value = update.Value
			if err := tx.Save(&setting).Error; err != nil {
				tx.Rollback()
				return c.Status(500).JSON(fiber.Map{"error": "Failed to update setting " + update.Key})
			}
		}
	}

	tx.Commit()
	return c.JSON(fiber.Map{"message": "Settings updated successfully"})
}
