package showtodo

import "github.com/kenfdev/cleancoding/server/adapter/gateway"

type showTodoInteractor struct {
	todoGateway gateway.TodoGateway
}

func NewShowTodoInteractor(tdGateway gateway.TodoGateway) ShowTodoInputPort {
	return &showTodoInteractor{
		todoGateway: tdGateway,
	}
}

func (it *showTodoInteractor) ShowTodo(requestModel *ShowTodoRequestModel, outputPort ShowTodoOutputPort) error {
	todo, err := it.todoGateway.FindTodoById(requestModel.ID)
	if err != nil {
		return err
	}

	res := &ShowTodoResponseModel{
		Todo: todo,
	}

	outputPort.Present(res)

	return nil
}
