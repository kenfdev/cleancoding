package todo

import (
	"github.com/kenfdev/cleancoding/server/adapter/controller"
	"github.com/kenfdev/cleancoding/server/usecase/listtodo"
)

type listTodoController struct {
	listTodoUseCase   listtodo.ListTodoInputPort
	listTodoPresenter listtodo.ListTodoOutputPort
}

func NewListTodoController(ltUseCase listtodo.ListTodoInputPort, ltPresenter listtodo.ListTodoOutputPort) *listTodoController {
	return &listTodoController{
		listTodoUseCase:   ltUseCase,
		listTodoPresenter: ltPresenter,
	}
}

func (c *listTodoController) Handle(ctx controller.Context) error {
	err := c.listTodoUseCase.ListTodo(c.listTodoPresenter)
	if err != nil {
		ctx.JSON(500, nil)
	}

	vm := c.listTodoPresenter.GetViewModel()
	ctx.JSON(200, vm)
	return nil
}
