package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Config struct {
	Address string `default: "localhost:8080"`
}

func main() {
	err := godotenv.Load()
	config := Config{
		Address: os.Getenv("ADDRESS"),
	}

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	err = app.Listen(config.Address)
	log.Fatal(err)
}
