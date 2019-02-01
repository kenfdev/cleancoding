package main

import (
	"github.com/kenfdev/echo-vue/server/infrastructure/api"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	api.ConfigureRouter(e)
	e.Logger.Fatal(e.Start(":1323"))
}
