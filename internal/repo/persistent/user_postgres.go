package persistent

import (
	"context"
	"fmt"

	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/repo"
	"github.com/London57/todo-app/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) CreateUser(ctx context.Context, user domain.User) (int, error) {
	usertable := repo.UserTable

	stmt, args, err := r.Builder.
		Insert(usertable.Name).
		Columns(usertable.Columns...).
		Values(user.Name, user.Username, user.Password).
		Suffix("returning id").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("UserRepo - CreateUser - r.Builder: %w", err)
	}

	var id int
	err = r.Pool.QueryRow(ctx, stmt, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("UserRepo - CreateUser - r.Pool.QueryRow: %w", err)
	}

	return id, nil
}