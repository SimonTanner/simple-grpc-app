package frontend

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Frontend struct {
	e *echo.Echo
}

func New() *Frontend {
	frontend := &Frontend{
		e: echo.New(),
	}

	frontend.e.GET("/", Hello)

	return frontend
}

func (f *Frontend) Start(address string) error {
	return f.e.Start(address)
}

func Hello(c echo.Context) error {
	response := Response{
		Msg:       "Hello You!",
		TimeStamp: time.Now(),
	}

	log.Println(response)

	return c.JSON(http.StatusOK, response)
}
