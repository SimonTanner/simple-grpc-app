package frontend

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
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
	frontend.e.GET("/booking", frontend.GetBooking)
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
	propertyParams := bookings.Property{}

	if c.QueryParam("country") != "" {
		log.Println(c.QueryParam("country"))
		propertyParams.Country = c.QueryParam("country")
	}

	if c.QueryParam("city") != "" {
		log.Println(c.QueryParam("city"))
		propertyParams.City = c.QueryParam("city")
	}

	log.Println("Getting properties from backend")

	stream, err := frontend.client.GetAllProperties(ctx, &propertyParams)

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
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

	log.Println("%+v\n", propertiesResult)

	return c.JSON(http.StatusOK, propertiesResult)
}

func (frontend *Frontend) CreateBooking(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	booking := BookingRequest{}

	err := c.Bind(&booking)
	if err != nil {
		log.Println(fmt.Sprintf("Error getting booking details from request: %v", err))
		return c.JSON(http.StatusBadRequest, err)
	}

	bookingMsg, err := booking.convertBookingToMsg()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	log.Println("Sending booking to backend")

	bookingCrtdMsg, err := frontend.client.BookPropertyById(ctx, bookingMsg)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating booking details: %v", err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	bookingResp, err := createBookingResponse(bookingCrtdMsg)
	if err != nil {
		log.Println(fmt.Sprintf("Error converting booking details: %v", err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, bookingResp)
}

func (frontend *Frontend) GetBooking(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if c.QueryParam("userId") == "" {
		return c.JSON(http.StatusBadRequest, errNoUserID)
	}

	log.Println(c.QueryParam("userId"))
	userID, err := strconv.Atoi(c.QueryParam("userId"))
	if err != nil {
		log.Printf("Incorrect type for userId param: %v", userID)
	}

	if c.QueryParam("propertyId") == "" {
		return c.JSON(http.StatusBadRequest, errNoPropertyID)
	}

	log.Println(c.QueryParam("userId"))
	propertyID, err := strconv.Atoi(c.QueryParam("propertyId"))
	if err != nil {
		log.Printf("Incorrect type for propertyId param: %v", propertyID)
	}

	if c.QueryParam("bookingId") == "" {
		return c.JSON(http.StatusBadRequest, errNoBookingID)
	}

	log.Println(c.QueryParam("bookingId"))
	bookingID, err := strconv.Atoi(c.QueryParam("bookingId"))
	if err != nil {
		log.Printf("Incorrect type for bookingId param: %v", bookingID)
		return c.JSON(http.StatusBadRequest, err)
	}

	upbMsg := createUPBMsg(userID, propertyID, bookingID)

	upbReceived, err := frontend.client.GetBookingsByBooking(ctx, upbMsg)
	if err != nil {
		log.Println(fmt.Sprintf("Error getting booking details from api: %v", err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	bkResp, err := createBookingResponse(upbReceived)
	if err != nil {
		log.Println(fmt.Sprintf("Error converting booking message to response: %v", err))
		return c.JSON(http.StatusInternalServerError, err)
	}

	if bkResp == (BookingResponse{}) {
		log.Println("No bookiings returned by backend api")
		return c.JSON(http.StatusInternalServerError, errNoBookings)
	}

	return c.JSON(http.StatusOK, bkResp)
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
