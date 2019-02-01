package controller

type Context interface {
	// Param returns path parameter by name.
	Param(name string) string
	// JSON sends a JSON response with status code.
	JSON(code int, i interface{}) error
}
