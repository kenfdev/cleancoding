package toggletodo

import (
	"github.com/kenfdev/cleancoding/server/domain"
)

type ToggleTodoRequestModel struct {
	ID         string
	IsComplete bool
}

type ToggleTodoResponseModel struct {
	Todo *domain.Todo
}

type ToggleTodoInputPort interface {
	ToggleTodo(requestModel *ToggleTodoRequestModel, outputPort ToggleTodoOutputPort) error
}

type ToggleTodoOutputPort interface {
	GetViewModel() *toggleTodoViewModel
	Present(model *ToggleTodoResponseModel)
}
