package database

import (
	"os"

	"github.com/gofiber/storage/postgres/v2"
)

var Store *postgres.Storage

func Connect() {
	Store = postgres.New(postgres.Config{
		ConnectionURI: os.Getenv("DATABASE_URL"),
	})
}
