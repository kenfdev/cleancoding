package todo

import (
	"github.com/kenfdev/echo-vue/server/adapter/controller"
	"github.com/kenfdev/echo-vue/server/domain"
	"github.com/kenfdev/echo-vue/server/usecase/createtodo"
)

type createTodoController struct {
	createTodoUseCase   createtodo.CreateTodoInputPort
	createTodoPresenter createtodo.CreateTodoOutputPort
}

func NewCreateTodoController(stUseCase createtodo.CreateTodoInputPort, stPresenter createtodo.CreateTodoOutputPort) *createTodoController {
	return &createTodoController{
		createTodoUseCase:   stUseCase,
		createTodoPresenter: stPresenter,
	}
}

func (c *createTodoController) Handle(ctx controller.Context) error {
	var m map[string]interface{}
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	// TODO: validation
	todo := &domain.Todo{
		Title:       m["title"].(string),
		Description: m["description"].(string),
		IsCompleted: false,
	}

	req := &createtodo.CreateTodoRequestModel{
		Todo: todo,
	}

	err := c.createTodoUseCase.CreateTodo(req, c.createTodoPresenter)
	if err != nil {
		ctx.JSON(500, err.Error())
		return err
	}

	vm := c.createTodoPresenter.GetViewModel()
	ctx.JSON(200, vm)

	return nil
}
