package gateway

import "github.com/kenfdev/cleancoding/server/domain"

type TodoGateway interface {
	FindAllTodos() ([]*domain.Todo, error)
	FindTodoById(id string) (*domain.Todo, error)
	ToggleTodoById(id string, isComplete bool) (*domain.Todo, error)
	CreateTodo(todo *domain.Todo) (*domain.Todo, error)
}
