package bookings

import (
	"time"

	pb "github.com/SimonTanner/simple-grpc-app/bookings"
	"github.com/golang/protobuf/ptypes"
)

type Property struct {
	Id         int       `db:"id"`
	DoorNumber string    `db:"door_number"`
	Address    string    `db:"address"`
	City       string    `db:"city"`
	Country    string    `db:"country"`
	CreatedAt  time.Time `db:"created_at"`
}

type User struct {
	Id        int       `db:"id"`
	FirstName string    `db:"first_name"`
	Surname   string    `db:"surname"`
	CreatedAt time.Time `db:"created_at"`
}

type Booking struct {
	PropertyId int       `db:"property_id"`
	UserId     int       `db:"user_id"`
	ID         int       `db:"user_id"`
	StartDate  time.Time `db:"start_date"`
	EndDate    time.Time `db:"end_date"`
	CreatedAt  time.Time `db:"created_at"`
}

type PropertyParams struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

func (user User) ConvertUserToMsg() (*pb.User, error) {
	createdAt, err := ptypes.TimestampProto(user.CreatedAt)
	if err != nil {
		return &pb.User{}, err
	}

	userMsg := pb.User{
		Id:        int32(user.Id),
		FirstName: user.FirstName,
		Surname:   user.Surname,
		CreatedAt: createdAt,
	}

	return &userMsg, nil
}

func (prop Property) ConvertPropertyToMsg() (*pb.Property, error) {
	createdAt, err := ptypes.TimestampProto(prop.CreatedAt)
	if err != nil {
		return &pb.Property{}, err
	}

	propertyMsg := pb.Property{
		Id:         int32(prop.Id),
		DoorNumber: prop.DoorNumber,
		Address:    prop.Address,
		City:       prop.City,
		Country:    prop.Country,
		CreatedAt:  createdAt,
	}

	return &propertyMsg, nil
}

func (b Booking) ConvertBookingToMsg() (*pb.Booking, error) {
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

	bookingMsg := pb.Booking{
		PropertyId: int32(b.PropertyId),
		UserId:     int32(b.UserId),
		StartDate:  startDate,
		EndDate:    endDate,
		CreatedAt:  createdAt,
	}

	return &bookingMsg, nil
}
