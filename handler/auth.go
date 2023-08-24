package handler

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alekangelov/testament/database"
	"github.com/alekangelov/testament/models"
	"github.com/alekangelov/testament/utils"
	"github.com/alekangelov/testament/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

// LoginHandler godoc
//
//	@Summary		Login
//	@Description	Login via email and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.LoginRequest	true	"Login"
//	@Success		200		{object}	models.LoginResponse
//	@Failure		400		{object}	models.GenericResponse
//	@Router			/auth/v1/auth/login [post]
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

	query := database.Store.Conn().QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", data.Email)

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
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid credentials",
		}
	}

	query = database.Store.Conn().QueryRow(context.Background(), "SELECT count(*) FROM verifications WHERE user_id = $1 AND type = $2", user.ID, "EMAIL_VERIFICATION")

	var verification models.CountModel

	if err := query.Scan(
		&verification.Count,
	); err != nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid credentials",
		}
	}

	if verification.Count > 0 {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "not verified",
		}
	}

	if !utils.ComparePassword(user.Password, data.Password) {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid credentials",
		}
	}

	token := utils.SignJWT(&utils.Claims{
		Subject:  user.ID,
		Iat:      time.Now().UTC().Unix(),
		Nbf:      time.Now().UTC().Unix(),
		Issuer:   "testament",
		Audience: []string{"testament"},
		Exp:      time.Now().UTC().Add(time.Hour * 24).Unix(),
	})

	refreshToken := utils.SignJWT(&utils.Claims{
		Subject:  user.ID,
		Iat:      time.Now().UTC().Unix(),
		Nbf:      time.Now().UTC().Unix(),
		Issuer:   "testament",
		Audience: []string{"testament"},
		Exp:      time.Now().UTC().Add(time.Hour * 24 * 7).Unix(),
	})

	return c.JSON(models.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	})
}

// RegisterHandler godoc
//
//	@Summary		Register
//	@Description	Register via email and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.RegisterRequest	true	"Register"
//	@Success		200		{object}	models.GenericResponse
//	@Failure		400		{object}	models.GenericResponse
//	@Router			/auth/v1/register [post]
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

// VerifyHandler godoc
//
//	@Summary		Verify
//	@Description	Verify some type of token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			body	body		models.VerifyRequest	true	"Verify"
//	@Success		200		{object}	models.GenericResponse
//	@Failure		400		{object}	models.GenericResponse
//	@Router			/auth/v1/verify [post]
func VerifyHandler(c *fiber.Ctx) error {
	data := new(models.VerifyRequest)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	if err := validator.Validator.Validate(data); len(err) > 0 {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "invalid code 1",
		}
	}

	query := database.Store.Conn().QueryRow(context.Background(), "SELECT * FROM verifications WHERE code = $1 AND type = $2 LIMIT 1", data.Code, data.Type)

	var verification models.VerificationModel

	if err := query.Scan(
		&verification.ID,
		&verification.UserId,
		&verification.Type,
		&verification.Code,
		&verification.CreatedAt,
		&verification.UpdatedAt,
	); err != nil {
		log.Printf("%+v", err)
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "verification not found",
		}
	}

	if row := database.Store.Conn().QueryRow(context.Background(), "DELETE FROM verifications WHERE id = $1 RETURNING *", verification.ID); row == nil {
		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "error deleting verification",
		}
	}

	return c.JSON(models.GenericResponse{
		Success: true,
		Message: "verified",
	})
}

// VerifyTokenHandler godoc
//
//	@Summary		Verify Token
//	@Description	Verify Token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.GenericResponse
//	@Failure		401	{object}	models.GenericResponse
//	@Security		Bearer
//
//	@Router			/auth/v1/verify-token [get]
func VerifyTokenHandler(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	token := strings.Split(authorization, " ")[1]
	_, err := utils.VerifyJWT(token)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	return c.JSON(models.GenericResponse{
		Success: true,
		Message: "verified",
	})
}

// RefreshTokenHandler godoc
//
//	@Summary		Refresh Token
//	@Description	Refresh Token
//	@Tags			auth
//	@Param			body	body	models.RefreshTokenRequest	true	"RefreshToken"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.LoginResponse
//	@Failure		401	{object}	models.GenericResponse
//	@Security		Bearer
//
//	@Router			/auth/v1/refresh-token [post]
func RefreshTokenHandler(c *fiber.Ctx) error {
	refreshToken := new(models.RefreshTokenRequest)

	if err := c.BodyParser(refreshToken); err != nil {
		return err
	}

	claims, err := utils.VerifyJWT(refreshToken.Token)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	if claims.Exp < time.Now().UTC().Unix() {
		return &fiber.Error{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
		}
	}

	token, err := utils.RefreshJWT(refreshToken.Token, 24)
	newRefresh, err := utils.RefreshJWT(refreshToken.Token, 24*7)

	return c.JSON(models.LoginResponse{
		Token:        token,
		RefreshToken: newRefresh,
	})
}
