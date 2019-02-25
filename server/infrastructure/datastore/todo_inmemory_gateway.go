package datastore

import (
	"errors"

	"github.com/kenfdev/cleancoding/server/adapter/gateway"
	"github.com/kenfdev/cleancoding/server/domain"
	uuid "github.com/satori/go.uuid"
)

var todos = []*domain.Todo{
	&domain.Todo{
		ID:          "1",
		Title:       "TODO TITLE",
		Description: "Todo Description",
		IsCompleted: false,
	},
	&domain.Todo{
		ID:          "2",
		Title:       "TODO TITLE2",
		Description: "Todo Description2",
		IsCompleted: false,
	},
	&domain.Todo{
		ID:          "3",
		Title:       "TODO TITLE3",
		Description: "Todo Description3",
		IsCompleted: false,
	},
}

type todoInMemoryGateway struct {
}

func NewTodoGateway() gateway.TodoGateway {
	return &todoInMemoryGateway{}
}

func (gw *todoInMemoryGateway) FindAllTodos() ([]*domain.Todo, error) {
	return todos, nil
}

func (gw *todoInMemoryGateway) FindTodoById(id string) (*domain.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, errors.New("Todo for ID: " + id + " not found.")
}

func (gw *todoInMemoryGateway) ToggleTodoById(id string, isComplete bool) (*domain.Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			todo.IsCompleted = isComplete
			return todo, nil
		}
	}
	return nil, errors.New("Todo for ID: " + id + " not found.")
}

func (gw *todoInMemoryGateway) CreateTodo(todo *domain.Todo) (*domain.Todo, error) {
	u1, err := uuid.NewV4()
	if err != nil {
		return nil, errors.New("UUID generation failed")
	}

	todo.ID = u1.String()

	todos = append(todos, todo)

	return todo, nil
}
