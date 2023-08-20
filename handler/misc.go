package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type HealthCheckResponse struct {
	Status string `json:"status"`
	Ok     bool   `json:"ok"`
	Now    int    `json:"now"`
}

func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(HealthCheckResponse{
		Status: "OK",
		Ok:     true,
		Now:    int(time.Now().UTC().UnixMilli()),
	})
}
