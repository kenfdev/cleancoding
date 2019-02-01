package api

import (
	"github.com/kenfdev/echo-vue/server/adapter/controller"
	"github.com/kenfdev/echo-vue/server/adapter/controller/todo"
	"github.com/kenfdev/echo-vue/server/infrastructure/datastore"
	"github.com/kenfdev/echo-vue/server/usecase/listtodo"
	"github.com/kenfdev/echo-vue/server/usecase/showtodo"
	"github.com/labstack/echo"
)

func passContext(fn func(c controller.Context) error) func(echo.Context) error {
	return func(c echo.Context) error {
		return fn(c)
	}
}

func ConfigureRouter(e *echo.Echo) {
	todoGateway := datastore.NewTodoGateway()
	listTodoInteractor := listtodo.NewListTodoInteractor(todoGateway)
	listTodoPresenter := listtodo.NewListTodoPresenter()
	listTodoController := todo.NewListTodoController(listTodoInteractor, listTodoPresenter)

	showTodoInteractor := showtodo.NewShowTodoInteractor(todoGateway)
	showTodoPresenter := showtodo.NewShowTodoPresenter()
	showTodoController := todo.NewShowTodoController(showTodoInteractor, showTodoPresenter)

	e.GET("/todos", passContext(listTodoController.Handle))
	e.GET("/todos/:id", passContext(showTodoController.Handle))
}
