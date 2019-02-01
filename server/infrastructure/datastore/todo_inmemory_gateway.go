package datastore

import (
	"errors"

	"github.com/kenfdev/echo-vue/server/adapter/gateway"
	"github.com/kenfdev/echo-vue/server/domain"
)

var todos = []*domain.Todo{
	&domain.Todo{
		ID:          "1",
		Title:       "TODO TITLE",
		Description: "Todo Description",
	},
	&domain.Todo{
		ID:          "2",
		Title:       "TODO TITLE2",
		Description: "Todo Description2",
	},
	&domain.Todo{
		ID:          "3",
		Title:       "TODO TITLE3",
		Description: "Todo Description3",
	},
}

type TodoInMemoryGateway struct {
}

func NewTodoGateway() gateway.TodoGateway {
	return &TodoInMemoryGateway{}
}

func (gw *TodoInMemoryGateway) FindAllTodos() ([]*domain.Todo, error) {
	return todos, nil
}

func (gw *TodoInMemoryGateway) FindTodoById(id string) (*domain.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, errors.New("Todo for ID: " + id + " not found.")
}
