package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email,max=255,min=3"`
		Password string `json:"password" validate:"required,max=255,min=3"`
	}

	RegisterRequest struct {
		Email                string `json:"email" validate:"required,email,max=255,min=3"`
		Password             string `json:"password" validate:"required,max=255,min=3,eqfield=PasswordConfirmation"`
		PasswordConfirmation string `json:"passwordConfirmation" validate:"required"`
	}

	ErrorResponse struct {
		Error       bool
		FailedField string
		Tag         string
		Value       interface{}
	}

	XValidator struct {
		Validator *validator.Validate
	}

	GlobalErrorHandlerResp struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	UserModel struct {
		ID             string    `json:"id"`
		Email          string    `json:"email"`
		Password       string    `json:"password"`
		SocialProvider *string   `json:"SOCIAL_PROVIDER"`
		SocialId       *string   `json:"SOCIAL_ID"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	VerificationModel struct {
		ID        string    `json:"id"`
		UserId    string    `json:"user_id"`
		Type      string    `json:"type"`
		Code      string    `json:"code"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
