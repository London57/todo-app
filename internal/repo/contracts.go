package repo

import (
	"context"

	"github.com/London57/todo-app/internal/entity"
)

type (
	UserRepo interface {
		CreateUser(context.Context, entity.User) (int, error)
	}
)
