package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/alekangelov/testament/database"
	"github.com/alekangelov/testament/models"
	"github.com/alekangelov/testament/utils"
	"github.com/alekangelov/testament/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
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

	hashedPassword := utils.HashPassword(data.Password)

	tx, err := database.Store.Conn().BeginTx(context.Background(), pgx.TxOptions{})

	if err != nil {
		return err
	}

	query := tx.QueryRow(context.Background(), "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *", data.Email, hashedPassword)

	var user models.UserModel

	if err := query.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.SocialProvider,
		&user.SocialId,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	query = tx.QueryRow(context.Background(), "INSERT INTO verifications (user_id, code, type) VALUES ($1, $2, $3) RETURNING *", user.ID, utils.GenerateVerificationCode(), "EMAIL_VERIFICATION")

	var verification models.VerificationModel

	if err := query.Scan(
		&verification.ID,
		&verification.UserId,
		&verification.Type,
		&verification.Code,
		&verification.CreatedAt,
		&verification.UpdatedAt,
	); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	if err := utils.SendMail([]string{user.Email}, fmt.Sprintf("Your verification code is: %s", verification.Code), "Welcome!"); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	if err := tx.Commit(context.Background()); err != nil {
		tx.Rollback(context.Background())
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "registered",
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
