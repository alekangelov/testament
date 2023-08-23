package handler

import (
	"fmt"
	"strings"

	"github.com/alekangelov/testament/models"
	"github.com/alekangelov/testament/validator"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(c *fiber.Ctx) error {
	data := new(models.LoginRequest)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	if err := validator.Validator.Validate(data); len(err) > 0 {
		errorMessages := make([]string, len(err))
		for i, err := range err {
			errorMessages[i] = fmt.Sprintf("%s", err.FailedField)
		}
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errorMessages, ", "),
		}
	}

	return c.JSON(fiber.Map{
		"token": "bearer",
	})
}

func RegisterHandler(c *fiber.Ctx) error {
	data := new(models.RegisterRequest)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	if err := validator.Validator.Validate(data); len(err) > 0 {
		errorMessages := make([]string, len(err))
		for i, err := range err {
			errorMessages[i] = fmt.Sprintf("%s", err.FailedField)
		}
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errorMessages, ", "),
		}
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func VerifyTokenHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func RefreshTokenHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
