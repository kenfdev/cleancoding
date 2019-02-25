package createtodo

import "github.com/kenfdev/cleancoding/server/adapter/gateway"

type createTodoInteractor struct {
	todoGateway gateway.TodoGateway
}

func NewCreateTodoInteractor(tdGateway gateway.TodoGateway) CreateTodoInputPort {
	return &createTodoInteractor{
		todoGateway: tdGateway,
	}
}

func (it *createTodoInteractor) CreateTodo(requestModel *CreateTodoRequestModel, outputPort CreateTodoOutputPort) error {
	todo, err := it.todoGateway.CreateTodo(requestModel.Todo)
	if err != nil {
		return err
	}

	res := &CreateTodoResponseModel{
		Todo: todo,
	}

	outputPort.Present(res)

	return nil
}
