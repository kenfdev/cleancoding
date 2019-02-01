package listtodo

import (
	"github.com/kenfdev/echo-vue/server/domain"
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
