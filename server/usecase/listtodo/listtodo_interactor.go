package listtodo

import "github.com/kenfdev/cleancoding/server/adapter/gateway"

type listTodoInteractor struct {
	todoGateway gateway.TodoGateway
}

func NewListTodoInteractor(tdGateway gateway.TodoGateway) ListTodoInputPort {
	return &listTodoInteractor{
		todoGateway: tdGateway,
	}
}

func (it *listTodoInteractor) ListTodo(outputPort ListTodoOutputPort) error {
	todos, err := it.todoGateway.FindAllTodos()
	if err != nil {
		return err
	}

	rm := &ListTodoResponseModel{
		Todos: todos,
	}
	outputPort.Present(rm)

	return nil
}
