package persistent

import "github.com/London57/todo-app/pkg/postgres"

type UserRepo struct {
	*postgres.Postgres
}
