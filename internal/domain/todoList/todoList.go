package todolist

import (
	"github.com/London57/todo-app/internal/domain"
	"github.com/London57/todo-app/internal/transport/todo_list"
)

type TodoList interface {
	CreateList(userID int, list todo_list.TodoListRequest) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId, listId int) (domain.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input domain.UpdateListInput) error
}
