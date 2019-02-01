package toggletodo

import "github.com/kenfdev/echo-vue/server/adapter/gateway"

type toggleTodoInteractor struct {
	todoGateway gateway.TodoGateway
}

func NewToggleTodoInteractor(tdGateway gateway.TodoGateway) ToggleTodoInputPort {
	return &toggleTodoInteractor{
		todoGateway: tdGateway,
	}
}

func (it *toggleTodoInteractor) ToggleTodo(requestModel *ToggleTodoRequestModel, outputPort ToggleTodoOutputPort) error {
	todo, err := it.todoGateway.ToggleTodoById(requestModel.ID, requestModel.IsComplete)
	if err != nil {
		return err
	}

	res := &ToggleTodoResponseModel{
		Todo: todo,
	}

	outputPort.Present(res)

	return nil
}
