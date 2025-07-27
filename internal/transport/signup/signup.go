package signup

import (
	"strings"
	"unicode/utf8"
)

type (
	SignUpResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}

	SignUpRequest struct {
		Name     string `json:"name" binding:"required,max=20,letteronly" validate:"required,max=20,letteronly"`
		Username string `json:"username" binding:"required,min=3,max=20,alphanum_underscore" validate:"required,min=3,max=20,alphanum_underscore"`
		Email    string `json:"email" binding:"required,email" validate:"required,email"`
		Password string `json:"password" binding:"required,max=20,min=6" validate:"required,max=20,min=6"`
	}

	SignUpOAuth2Request struct {
		Name  string `json:"name" binding:"required,max=20,letteronly" validate:"required,max=20,letteronly"`
		Email string `json:"email" binding:"required,email" validate:"required,email"`
	}

	SignInRequest struct {
		UsernameOrEmail string `json:"username_or_email" binding:"min=3,required" validate:"min=3,required"`
		Password        string `json:"password" binding:"required,min=6,max=20" validate:"required,min=6,max=20"`
	}
)

func IsEmail(str string) bool {
	ai := strings.Index(str, "@")
	di := strings.Index(str, ".")
	if ai > 0 && di > ai + 1 && di < utf8.RuneCountInString(str) - 1 {
		return true
	}
	return false
}