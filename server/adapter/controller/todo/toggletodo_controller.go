package todo

import (
	"errors"

	"github.com/kenfdev/echo-vue/server/adapter/controller"
	"github.com/kenfdev/echo-vue/server/usecase/toggletodo"
)

type toggleTodoController struct {
	toggleTodoUseCase   toggletodo.ToggleTodoInputPort
	toggleTodoPresenter toggletodo.ToggleTodoOutputPort
}

func NewToggleTodoController(stUseCase toggletodo.ToggleTodoInputPort, stPresenter toggletodo.ToggleTodoOutputPort) *toggleTodoController {
	return &toggleTodoController{
		toggleTodoUseCase:   stUseCase,
		toggleTodoPresenter: stPresenter,
	}
}

func (c *toggleTodoController) Handle(ctx controller.Context) error {
	id := ctx.Param("id")
	if id == "" {
		msg := "Bad Request"
		ctx.JSON(400, msg)
		return errors.New(msg)
	}

	var m map[string]interface{}
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	req := &toggletodo.ToggleTodoRequestModel{
		ID:         id,
		IsComplete: m["isCompleted"].(bool),
	}

	err := c.toggleTodoUseCase.ToggleTodo(req, c.toggleTodoPresenter)
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}

	vm := c.toggleTodoPresenter.GetViewModel()
	ctx.JSON(200, vm)

	return nil
}
