package v1

import (
	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/v1/auth"
	"github.com/London57/todo-app/internal/controller/http/v1/item"
	"github.com/London57/todo-app/internal/controller/http/v1/list"
)

type V1 struct {
	Auth auth.AuthController
	List list.ListController
	Item item.ItemController
	env  *config.Config
}

func New(a auth.AuthController, l list.ListController, i item.ItemController, env *config.Config) V1 {
	return V1{
		Auth: a,
		List: l,
		Item: i,
		env:  env,
	}
}
