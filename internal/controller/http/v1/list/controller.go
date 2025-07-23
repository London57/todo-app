package list

import "github.com/London57/todo-app/internal/controller/http/common"

type ListController struct {
	common.BaseController
	// uc
}

func New(bC common.BaseController) ListController {
	return ListController{
		BaseController: bC,
	}
}
