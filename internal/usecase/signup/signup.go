package signup

import (
	"context"
	"fmt"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/domain/jwtutil"
	"github.com/London57/todo-app/internal/domain/password"
	"github.com/London57/todo-app/internal/domain/signup"
	"github.com/London57/todo-app/internal/repo"
)

type SignUpUseCase struct {
	repo repo.UserRepo
}

func New(r repo.UserRepo) *SignUpUseCase {
	return &SignUpUseCase{
		repo: r,
	}
}

func (uc *SignUpUseCase) CreateUser(context context.Context, user signup.SignUpRequest) (int, error) {
	var err error
	if user.Password != "" {
		user.Password, err = password.GeneratePasswordHash(user.Password)
		if err != nil {
			return 0, err
		}
	}

	id, err := uc.repo.CreateUser(context, user)
	if err != nil {
		return 0, fmt.Errorf("UserUseCase - CreateUser - repo.CreateUser: %w", err)
	}
	return id, nil
}

func (uc *SignUpUseCase) GetUserByUsername(context context.Context, username string) (domain.User, error) {
	user, err := uc.repo.GetUserByUsername(context, username)
	if err != nil {
		return domain.User{}, ErrUserNotFound
	}
	return user, nil
}

func (uc *SignUpUseCase) CreateAccessToken(user domain.User, secret string, expiry int) (string, error) {
	token, err := jwtutil.CreateAccessToken(user, secret, expiry)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *SignUpUseCase) CreateRefreshToken(user domain.User, secret string, expiry int) (string, error) {
	token, err := jwtutil.CreateRefreshToken(user, secret, expiry)
	if err != nil {
		return "", err
	}
	return token, nil
}
