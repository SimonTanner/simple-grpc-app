package frontend

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "github.com/SimonTanner/simple-grpc-app/bookings"
)

type Response struct {
	Msg       string    `json:"msg"`
	TimeStamp time.Time `json:"timeStamp"`
}

type BookingRequest struct {
	PropertyId int32     `json:"propertyId"`
	UserId     int32     `json:"userId"`
	StartDate  time.Time `json:"startDate"`
	EndDate    time.Time `json:"endDate"`
	CreatedAt  time.Time `json:"createdAt"`
}

type BookingResponse struct {
	Property  Property  `json:"property"`
	User      User      `json:"user"`
	StartDate time.Time `json:"startDate:omitempty"`
	EndDate   time.Time `json:"endDate:omitempty"`
	CreatedAt time.Time `json:"createdAt:omitempty"`
}

type Property struct {
	Id         int32  `json:"id"`
	DoorNumber string `json:"doorNumber"`
	Address    string `json:"address"`
	City       string `json:"city"`
	Country    string `json:"country"`
}

type User struct {
	Id        int32  `json:"id"`
	FirstName string `json:"firstName"`
	Surname   string `json:"surname"`
}

func (b BookingRequest) convertBookingToMsg() (*pb.Booking, error) {
	startDate, err := ptypes.TimestampProto(b.StartDate)
	if err != nil {
		return &pb.Booking{}, err
	}

	endDate, err := ptypes.TimestampProto(b.EndDate)
	if err != nil {
		return &pb.Booking{}, err
	}

	createdAt, err := ptypes.TimestampProto(b.CreatedAt)
	if err != nil {
		return &pb.Booking{}, err
	}

	// var bookingMsg pb.Booking

	bookingMsg := pb.Booking{
		PropertyId: b.PropertyId,
		UserId:     b.UserId,
		StartDate:  startDate,
		EndDate:    endDate,
		CreatedAt:  createdAt,
	}

	return &bookingMsg, nil
}

func createBookingResponse(bk *pb.UserPropertyBooking) (BookingResponse, error) {
	user := User{
		Id:        bk.User.Id,
		FirstName: bk.User.FirstName,
		Surname:   bk.User.Surname,
	}

	property := Property{
		Id:         bk.Property.Id,
		DoorNumber: bk.Property.DoorNumber,
		Address:    bk.Property.Address,
		City:       bk.Property.City,
		Country:    bk.Property.Country,
	}

	booking := BookingResponse{
		Property: property,
		User: user,s
	}

	return booking, nil
}
