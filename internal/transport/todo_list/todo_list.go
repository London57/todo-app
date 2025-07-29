package todo_list

type TodoListRequest struct {
	Title       string `json:"title" binding:"required" validate:"required"`
	Description string `json:"description"`
}
