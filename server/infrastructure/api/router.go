package api

import (
	"github.com/kenfdev/cleancoding/server/adapter/controller"
	"github.com/kenfdev/cleancoding/server/adapter/controller/todo"
	"github.com/kenfdev/cleancoding/server/infrastructure/datastore"
	"github.com/kenfdev/cleancoding/server/usecase/createtodo"
	"github.com/kenfdev/cleancoding/server/usecase/listtodo"
	"github.com/kenfdev/cleancoding/server/usecase/showtodo"
	"github.com/kenfdev/cleancoding/server/usecase/toggletodo"
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

	toggleTodoInteractor := toggletodo.NewToggleTodoInteractor(todoGateway)
	toggleTodoPresenter := toggletodo.NewToggleTodoPresenter()
	toggleTodoController := todo.NewToggleTodoController(toggleTodoInteractor, toggleTodoPresenter)

	createTodoInteractor := createtodo.NewCreateTodoInteractor(todoGateway)
	createTodoPresenter := createtodo.NewCreateTodoPresenter()
	createTodoController := todo.NewCreateTodoController(createTodoInteractor, createTodoPresenter)

	e.GET("/api/todos", passContext(listTodoController.Handle))
	e.GET("/api/todos/:id", passContext(showTodoController.Handle))
	e.PUT("/api/todos/:id/complete", passContext(toggleTodoController.Handle))
	e.POST("/api/todos", passContext(createTodoController.Handle))
}
