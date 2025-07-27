package validate

import "github.com/go-playground/validator/v10"

func NewValidator() *validator.Validate {
	validator := validator.New()
	validator.RegisterValidation("alphanum_underscore", AlphanumAndUnderscoreValidate)
	validator.RegisterValidation("letteronly", LetterValidate)

	return validator
}