package usecase

import (
	"context"

	"github.com/London57/todo-app/internal/entity"
)

type (
	User interface {
		CreateUser(context.Context, entity.User) (int, error)
	}
)
