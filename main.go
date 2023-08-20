package main

import (
	"log"
	"os"

	"github.com/alekangelov/testament/database"
	"github.com/alekangelov/testament/models"
	"github.com/alekangelov/testament/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

type Config struct {
	Address string `default: "localhost:8080"`
}

func main() {

	err := godotenv.Load()

	database.Connect()

	database.Migrate()

	config := Config{
		Address: os.Getenv("ADDRESS"),
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(models.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
	app.Use(cors.New())

	router.SetupRoutes(app)

	err = app.Listen(config.Address)
	log.Fatal(err)
}
