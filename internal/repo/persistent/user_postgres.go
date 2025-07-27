package persistent

import (
	"context"
	"fmt"
	
	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/transport/signup"
	"github.com/London57/todo-app/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type UserRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) CreateUser(ctx context.Context, user signup.SignUpRequest) (uuid.UUID, error) {
	stmt, args, err := r.Builder.
		Insert("\"user\"").
		Columns("name", "username", "email", "password").
		Values(user.Name, user.Username, user.Email, user.Password).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("UserRepo - CreateUser - r.Builder: %w", err)
	}

	var id uuid.UUID
	err = r.Pool.QueryRow(ctx, stmt, args...).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("UserRepo - CreateUser - r.Pool.QueryRow: %w", err)
	}

	return id, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	stmt, args, err := r.Builder.
		Select("*").
		From("\"user\"").
		Where(squirrel.Eq{"email": email}).
		ToSql()
	if err != nil {
		return domain.User{}, err
	}

	var user domain.User
	err = r.Pool.QueryRow(ctx, stmt, args...).Scan(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepo) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	stmt, args, err := r.Builder.Select("*").
    From("\"user\"").
		Where(squirrel.Eq{"username": username}).
		ToSql()
	if err != nil {
		return domain.User{}, err
	}
	var user domain.User
	err = r.Pool.QueryRow(ctx, stmt, args).Scan(&user)
	if err != nil {
		return domain.User{}, nil
	}
	return user, nil
}
