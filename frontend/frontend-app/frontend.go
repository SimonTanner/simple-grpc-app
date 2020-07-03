package frontend

import (
	"context"
	"fmt"
	"io"
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
	frontend.e.POST("/booking", frontend.CreateBooking)

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

	propertiesResult := []*bookings.Property{}

	for {
		prop, err := stream.Recv()
		log.Println(prop)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Error receiving properties from stream %v", err)
		}

		propertiesResult = append(propertiesResult, prop)
	}

	fmt.Printf("%+v\n", propertiesResult)

	return c.JSON(http.StatusOK, propertiesResult)
}

func (frontend *Frontend) CreateBooking(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var booking BookingRequest

	err := c.Bind(booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingMsg, err := booking.convertBookingToMsg()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	log.Println("Sending booking to backend")

	bookingCrtdMsg, err := frontend.client.BookPropertyById(ctx, bookingMsg)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating booking details: %v", err))
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingResp, err := createBookingResponse(bookingCrtdMsg)
	if err != nil {
		log.Println(fmt.Sprintf("Error converting booking details: %v", err))
		return c.JSON(http.StatusInternalServerErrors, err)
	}

	return c.JSON(http.StatusCreated, bookingResp)
}

// func get() {
// 	userId, err := c.Param("userId")
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	prop, err := c.Param("userId")
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}

// 	userId, err := c.Param("userId")
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err)
// 	}
// }
