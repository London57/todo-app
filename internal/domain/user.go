package domain

import (
	"context"

	"github.com/google/uuid"
)

type (
	User struct {
		ID       uuid.UUID `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	UserRequest struct {
		Name     string `json:"name" validate:"required,max=20"`
		Username string `json:"username" validate:"required,max=20"`
		Password string `json:"password" validate:"required,max=15,min=9"`
	}
)

type (
	SignUpUseCase interface {
		RegisterUserByUsername(context.Context, User) (int, error)
		RegisterUserByOAuth2(context.Context) (int, error)
	}
)