package signup

import (
	"context"

	"github.com/London57/todo-app/internal/transport/signup"
	"github.com/London57/todo-app/internal/domain"
	"github.com/google/uuid"
)

//go:generate mockgen -source=signup.go -destination=mocks/mock.go

type SignUpUseCase interface {
	CreateUser(context.Context, signup.SignUpRequest) (uuid.UUID, error)
	GetUserByEmail(context.Context, string) (domain.User, error)
	GetUserByUsername(context.Context, string) (domain.User, error)
	CreateAccessToken(domain.User, string, int) (string, error)
	CreateRefreshToken(domain.User, string, int) (string, error)
}
