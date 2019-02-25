package listtodo

import (
	"github.com/kenfdev/cleancoding/server/domain"
)

type ListTodoResponseModel struct {
	Todos []*domain.Todo
}

type ListTodoInputPort interface {
	ListTodo(outputPort ListTodoOutputPort) error
}

type ListTodoOutputPort interface {
	GetViewModel() *listTodoViewModel
	Present(model *ListTodoResponseModel)
}
