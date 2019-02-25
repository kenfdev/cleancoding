package createtodo

import (
	"github.com/kenfdev/cleancoding/server/domain"
)

type CreateTodoRequestModel struct {
	Todo *domain.Todo
}

type CreateTodoResponseModel struct {
	Todo *domain.Todo
}

type CreateTodoInputPort interface {
	CreateTodo(requestModel *CreateTodoRequestModel, outputPort CreateTodoOutputPort) error
}

type CreateTodoOutputPort interface {
	GetViewModel() *createTodoViewModel
	Present(model *CreateTodoResponseModel)
}
