package handler

import (
	"encoding/json"
	"os"

	"github.com/alekangelov/testament/models"
	"github.com/gofiber/fiber/v2"
)

func SettingsHandler(c *fiber.Ctx) error {
	settingsJson := os.Getenv("SETTINGS")
	settings := models.SettingsResponse{}
	if settingsJson != "" {
		json.Unmarshal([]byte(settingsJson), &settings)
	}
	return c.JSON(settings)
}
