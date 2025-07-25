package signup

import (
	"context"
	"fmt"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/domain/jwtutil"
	"github.com/London57/todo-app/internal/domain/password"
	"github.com/London57/todo-app/internal/domain/signup"
	"github.com/London57/todo-app/internal/repo"
	"github.com/google/uuid"
)

type SignUpUseCase struct {
	repo repo.UserRepo
}

func New(r repo.UserRepo) SignUpUseCase {
	return SignUpUseCase{
		repo: r,
	}
}

func (uc *SignUpUseCase) CreateUser(context context.Context, user signup.SignUpRequest) (uuid.UUID, error) {
	var err error
	if user.Password != "" {
		user.Password, err = password.GeneratePasswordHash(user.Password)
		if err != nil {
			return uuid.Nil, err
		}
	}

	id, err := uc.repo.CreateUser(context, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("UserUseCase - CreateUser - repo.CreateUser: %w", err)
	}
	return id, nil
}

func (uc *SignUpUseCase) GetUserByEmail(context context.Context, email string) (domain.User, error) {
	user, err := uc.repo.GetUserByEmail(context, email)
	if err != nil {
		return domain.User{}, err
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
