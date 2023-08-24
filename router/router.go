package router

import (
	"github.com/alekangelov/testament/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	app.Get("/health", handler.HealthCheck)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Type("HTML").SendString("<h1>Testament</h1>")
	})

	auth := api.Group("v1/auth")
	auth.Post("/login", handler.LoginHandler)
	auth.Post("/register", handler.RegisterHandler)
	auth.Post("/verify", handler.VerifyHandler)
	auth.Get("/verify-token", handler.VerifyTokenHandler)
	auth.Post("/refresh-token", handler.RefreshTokenHandler)
}
