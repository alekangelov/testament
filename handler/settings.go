package handler

import (
	"github.com/alekangelov/testament/models"
	"github.com/gofiber/fiber/v2"
)

func SettingsHandler(c *fiber.Ctx) error {
	return c.JSON(models.SettingsResponse{
		Title:       "Testament",
		Description: "Testament is a simple, fast and secure Authentication server.",
	})
}
