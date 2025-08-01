package todo_list

import "errors"

type TodoListRequest struct {
	Title       string `json:"title" binding:"required" validate:"required"`
	Description string `json:"description"`
}

type UpdateListRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (i UpdateListRequest) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
