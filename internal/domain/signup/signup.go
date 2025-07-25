package signup

import (
	"context"

	"github.com/London57/todo-app/internal/domain"
	"github.com/google/uuid"
)

type (
	SignUpRequest struct {
		Name     string `json:"name" validate:"required,max=20"`
		Username string `json:"username" validate:"required,max=20,al"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,max=15,min=9"`
	}
	SignUpOAuth2Request struct {
		Name  string `json:"name" validate:"required,max=20"`
		Email string `json:"email" validate:"required,email"`
	}
	SignUpResponse struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}
)

type (
	SignUpUseCase interface {
		CreateUser(context.Context, SignUpRequest) (uuid.UUID, error)
		GetUserByEmail(context.Context, string) (domain.User, error)
		CreateAccessToken(domain.User, string, int) (string, error)
		CreateRefreshToken(domain.User, string, int) (string, error)
	}
)
