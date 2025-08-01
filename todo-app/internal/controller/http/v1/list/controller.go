package list

import (
	"github.com/London57/todo-app/internal/controller/http/common/controller"
)

type ListController struct {
	controller.BaseController
}

func NewListController(bC controller.BaseController) ListController {
	return ListController{
		BaseController: bC,
	}
}
