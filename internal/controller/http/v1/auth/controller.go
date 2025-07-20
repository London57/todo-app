package auth

import (
	"github.com/London57/todo-app/internal/controller/http/common"
	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/infra/alias"
)

type AuthController struct {
	common.BaseController
	domain.SignUpUseCase
	extra_data alias.Extra_data
}

func NewAuthController(b common.BaseController, sauc domain.SignUpUseCase, extra_d alias.Extra_data) *AuthController {
	return &AuthController{b, sauc, extra_d}
}