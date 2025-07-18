package auth

import (
	"context"
	"fmt"

	"github.com/London57/todo-app/internal/entity"
	"github.com/London57/todo-app/internal/repo"
)

type UseCase struct {
	repo repo.UserRepo
}

func New(r repo.UserRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

func (uc *UseCase) CreateUser(context context.Context, user entity.User) (int, error) {
	id, err := uc.repo.CreateUser(context, user)
	if err != nil {
		return 0, fmt.Errorf("UserUseCase - CreateUser - repo.CreateUser: %w", err)
	}
	return id, nil
}
