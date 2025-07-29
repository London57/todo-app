package item

import (
	"net/http"

	er "github.com/London57/todo-app/internal/controller/http/error"
	"github.com/London57/todo-app/internal/controller/http/middleware"
	"github.com/London57/todo-app/internal/transport/todo_list"
	"github.com/gin-gonic/gin"
)

func (h *ItemController) CreateItem(r *gin.Context) {
	id, ok := r.Get(middleware.UserID)
	if !ok {
		er.ErrorResponse(r, http.StatusInternalServerError, "user id not found", "")
		return
	}
	var input todo_list.TodoListRequest
	if err := r.BindJSON(&input); err != nil {
		er.ErrorResponse(r, http.StatusBadRequest, "failed to parse ", "")
		return
	}

}

func (h *ItemController) GetAllItems(r *gin.Context) {}

func (h *ItemController) GetItemById(r *gin.Context) {}

func (h *ItemController) UpdateItem(r *gin.Context) {}

func (h *ItemController) DeleteItem(r *gin.Context) {}
