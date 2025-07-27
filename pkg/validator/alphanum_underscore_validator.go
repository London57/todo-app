package validate

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

func AlphanumAndUnderscoreValidate(fl validator.FieldLevel) bool {
	for _, char := range fl.Field().String() {
		if !unicode.In(char, unicode.Latin) && char != '_' && !unicode.IsDigit(char){
			return false
		}
	}
	return true
}