package auth

import (
	"github.com/London57/todo-app/config"
	"github.com/London57/todo-app/internal/controller/http/common/controller"
	"github.com/London57/todo-app/internal/domain/signup"
)

type AuthController struct {
	controller.BaseController
	signup.SignUpUseCase
	env *config.Config
}

func NewAuthController(b controller.BaseController, sauc signup.SignUpUseCase, env *config.Config) AuthController {
	return AuthController{b, sauc, env}
}
