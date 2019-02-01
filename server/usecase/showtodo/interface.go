package showtodo

import (
	"github.com/kenfdev/echo-vue/server/domain"
)

type ShowTodoRequestModel struct {
	ID string
}

type ShowTodoResponseModel struct {
	Todo *domain.Todo
}

type ShowTodoInputPort interface {
	ShowTodo(requestModel *ShowTodoRequestModel, outputPort ShowTodoOutputPort) error
}

type ShowTodoOutputPort interface {
	GetViewModel() *showTodoViewModel
	Present(model *ShowTodoResponseModel)
}
