package auth

import (
	"context"
	"fmt"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/infra/alias"
	"github.com/London57/todo-app/internal/infra/jwtutil"
	"github.com/London57/todo-app/internal/repo"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	repo repo.UserRepo
	extra_data alias.Extra_data
}

func New(r repo.UserRepo, extra_data alias.Extra_data) *UseCase {
	return &UseCase{
		repo: r,
		extra_data: extra_data,
	}
}

func (uc *UseCase) RegisterUserByUsername(context context.Context, user domain.User) (int, error) {
	var err error
	user.Password, err = generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	
	id, err := uc.repo.CreateUser(context, user)
	if err != nil {
		return 0, fmt.Errorf("UserUseCase - CreateUser - repo.CreateUser: %w", err)
	}
	x = jwtutil.CreateAccessToken()
	return id, nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("UserUsecase - generatePasswordHash: %w", err)
	}
	return string(hash), nil
}

