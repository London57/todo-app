package repo

import (
	"context"

	"github.com/London57/todo-app/internal/domain"
)

type (
	UserRepo interface {
		CreateUser(context.Context, domain.User) (int, error)
		GetUserByUsername(context.Context, string) (domain.User, error)
	}
)
