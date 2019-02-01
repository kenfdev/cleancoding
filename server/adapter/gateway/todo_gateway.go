package gateway

import "github.com/kenfdev/echo-vue/server/domain"

type TodoGateway interface {
	FindAllTodos() ([]*domain.Todo, error)
	FindTodoById(id string) (*domain.Todo, error)
}
