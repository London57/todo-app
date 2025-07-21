package v1

import (
	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/v1/auth"
	"github.com/London57/todo-app/internal/controller/http/v1/item"
	"github.com/London57/todo-app/internal/controller/http/v1/list"
)

type V1 struct {
	Auth *auth.AuthController
	List *list.ListController
	Item *item.ItemController
	env  config.Config
}
