package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {
	queryValue, err := strconv.Atoi(c.Query("time"))
	if err != nil {
		queryValue = int(time.Now().UTC().UnixMilli())
	}

	now := int(time.Now().UTC().UnixMilli())

	return c.JSON(fiber.Map{
		"status": "ok",
		"ok":     true,
		"now":    now,
		"diff":   now - queryValue,
	})
}
