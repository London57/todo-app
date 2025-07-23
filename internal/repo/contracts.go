package repo

import (
	"context"

	"github.com/London57/todo-app/internal/domain"
	"github.com/google/uuid"
)

type (
	UserRepo interface {
		CreateUser(context.Context, domain.User) (uuid.UUID, error)
		GetUserByEmail(context.Context, string) (domain.User, error)
	}
)
