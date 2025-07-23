package item

import (
	"github.com/London57/todo-app/internal/controller/http/common"
)

type ItemController struct {
	common.BaseController
	// uc
}

func New(bC common.BaseController) ItemController {
	return ItemController{
		BaseController: bC,
	}
}
