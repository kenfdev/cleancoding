package todo

import (
	"errors"

	"github.com/kenfdev/cleancoding/server/adapter/controller"
	"github.com/kenfdev/cleancoding/server/usecase/showtodo"
)

type showTodoController struct {
	showTodoUseCase   showtodo.ShowTodoInputPort
	showTodoPresenter showtodo.ShowTodoOutputPort
}

func NewShowTodoController(stUseCase showtodo.ShowTodoInputPort, stPresenter showtodo.ShowTodoOutputPort) *showTodoController {
	return &showTodoController{
		showTodoUseCase:   stUseCase,
		showTodoPresenter: stPresenter,
	}
}

func (c *showTodoController) Handle(ctx controller.Context) error {
	id := ctx.Param("id")

	if id == "" {
		msg := "Bad Request"
		ctx.JSON(400, msg)
		return errors.New(msg)
	}

	req := &showtodo.ShowTodoRequestModel{
		ID: id,
	}

	err := c.showTodoUseCase.ShowTodo(req, c.showTodoPresenter)
	if err != nil {
		ctx.JSON(500, nil)
		return errors.New("")
	}

	vm := c.showTodoPresenter.GetViewModel()
	ctx.JSON(200, vm)

	return nil
}
