package validator

import (
	"github.com/alekangelov/testament/models"
	"github.com/go-playground/validator/v10"
)

type (
	XValidator struct {
		Validator *validator.Validate
	}
)

var validate = validator.New()

func (v *XValidator) Validate(i interface{}) []models.ErrorResponse {
	validationErrors := []models.ErrorResponse{}

	err := v.Validator.Struct(i)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, models.ErrorResponse{
				Error:       true,
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Param(),
			})
		}
	}

	return validationErrors
}

var Validator = &XValidator{Validator: validate}
