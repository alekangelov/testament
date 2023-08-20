package router

import (
	"github.com/alekangelov/testament/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	app.Get("/health", handler.HealthCheck)

	auth := api.Group("/auth")
	auth.Post("/login", handler.LoginHandler)
	auth.Post("/register", handler.RegisterHandler)
	auth.Post("/verify", handler.VerifyHandler)
}
