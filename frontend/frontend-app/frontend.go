package frontend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SimonTanner/simple-grpc-app/bookings"
	"github.com/labstack/echo/v4"
)

type Frontend struct {
	e      *echo.Echo
	client bookings.BookingsApiClient
}

func New(client bookings.BookingsApiClient) *Frontend {
	frontend := Frontend{
		e:      echo.New(),
		client: client,
	}

	frontend.e.GET("/", Hello)
	frontend.e.GET("/properties", frontend.GetProperties)

	return &frontend
}

func (f *Frontend) Start(address string) error {
	return f.e.Start(address)
}

func Hello(c echo.Context) error {
	response := Response{
		Msg:       "Welcome to Book My Place!",
		TimeStamp: time.Now(),
	}

	log.Println(response)

	return c.JSON(http.StatusOK, response)
}

func (frontend *Frontend) GetProperties(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Getting properties from backend")

	stream, err := frontend.client.GetAllProperties(ctx, &bookings.Property{})

	if err != nil {
		log.Println(err)
		return err
	}

	props, _ := stream.Recv()

	log.Println(props)
	// for _, prop := range props {

	// }
	fmt.Printf("%+v\n", props)

	return c.JSON(http.StatusOK, props)
}
