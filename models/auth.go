package models

import (
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
)
