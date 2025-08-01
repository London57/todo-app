package validate

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func LetterValidate(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}