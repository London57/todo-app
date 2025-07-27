package validate

import "github.com/go-playground/validator/v10"

type CustomValidator struct {
	V *validator.Validate
}

func (v *CustomValidator) ValidateStruct(obj any) error {
	return v.V.Struct(obj)
}

func (v *CustomValidator) Engine() any {
	return v.V
}

func NewValidator() *validator.Validate {
	validator := validator.New()
	validator.RegisterValidation("alphanum_underscore", AlphanumAndUnderscoreValidate)
	validator.RegisterValidation("letteronly", LetterValidate)

	return validator
}
