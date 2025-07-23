package common

import (
	"github.com/London57/todo-app/pkg/logger"
	"github.com/go-playground/validator/v10"
)

type BaseController struct {
	L logger.Interface
	V *validator.Validate
}

func New(l logger.Interface, v *validator.Validate) BaseController {
	return BaseController{
		L: l,
		V: v,
	}
}
