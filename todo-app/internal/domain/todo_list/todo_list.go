package todo_list

import (
	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/transport/todo_list"
)

type TodoListUseCase interface {
	CreateList(userID int, list todo_list.TodoListRequest) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId, listId int) (domain.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo_list.UpdateListRequest) error
}
