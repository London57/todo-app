package repo

import (
	"context"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/transport/signup"
	"github.com/London57/todo-app/internal/transport/todo_list"
	"github.com/google/uuid"
)

type (
	UserRepo interface {
		Create(context.Context, signup.SignUpRequest) (uuid.UUID, error)
		GetByEmail(context.Context, string) (domain.User, error)
		GetByUsername(context.Context, string) (domain.User, error)
	}

	TodoListRepo interface {
		Create(context context.Context, userID uuid.UUID, list todo_list.TodoListRequest) (int, error)
		GetAll(context context.Context, userID uuid.UUID) ([]domain.TodoList, error)
		GetById(context context.Context, userID uuid.UUID, listID int) (domain.TodoList, error)
		Delete(context context.Context, userID uuid.UUID, listID int) error
		Update(context context.Context, userID uuid.UUID, input todo_list.UpdateListRequest) error
	}
)
