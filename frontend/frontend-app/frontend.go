package frontend

import "github.com/labstack/echo/v4"

type Frontend struct {
	e *echo.Echo
}

func New() *Frontend {
	frontend := &Frontend{
		e: echo.New(),
	}

	return frontend
}
