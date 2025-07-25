package repo

import (
	"context"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/domain/signup"
	"github.com/google/uuid"
)

type (
	UserRepo interface {
		CreateUser(context.Context, signup.SignUpRequest) (uuid.UUID, error)
		GetUserByEmail(context.Context, string) (domain.User, error)
	}
)
