package item

import (
	"github.com/London57/todo-app/internal/controller/http/common/controller"
)

type ItemController struct {
	controller.BaseController
}

func NewItemController(bC controller.BaseController) ItemController {
	return ItemController{
		BaseController: bC,
	}
}
