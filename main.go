package main

import (
	"log"
	"os"

	"github.com/alekangelov/testament/database"
	_ "github.com/alekangelov/testament/docs"
	"github.com/alekangelov/testament/models"
	"github.com/alekangelov/testament/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
)

//	@title			Testament API
//	@version		1.0
//	@description	This is the API for Testament

//	@license.name	MIT
//	@license.url	/license.md

//	@host		localhost:8080
//	@BasePath	/

func main() {

	err := godotenv.Load()

	database.Connect()

	database.Migrate()

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = ":8080"
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(models.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	app.Use(cors.New(
		cors.Config{
			AllowOrigins:     "*",
			AllowHeaders:     "*",
			AllowMethods:     "*",
			AllowCredentials: true,
		},
	))
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/license", func(c *fiber.Ctx) error {
		return c.SendFile("./license.md")
	})

	router.SetupRoutes(app)

	err = app.Listen(address)
	log.Fatal(err)
}
