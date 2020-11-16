package api

import (
	"context"
	"fmt"
	"log"

	"github.com/SimonTanner/simple-grpc-app/backend/bookings"
	pb "github.com/SimonTanner/simple-grpc-app/bookings"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	pb.UnimplementedBookingsApiServer
	dbService bookings.Service
	// properties []*pb.Property
}

func New(db *sqlx.DB) *Api {

	api := &Api{
		dbService: bookings.New(db),
	}
	return api
}

func (a *Api) GetAllProperties(p *pb.Property, stream pb.BookingsApi_GetAllPropertiesServer) error {

	log.Println(p)
	log.Println("Getting properties")

	params := bookings.PropertyParams{
		City:    p.City,
		Country: p.Country,
	}

	properties, err := a.dbService.GetAllProperties(params)

	fmt.Println(properties)

	if err != nil {
		log.Println(err)
		return err
	}

	grpcProps := []*pb.Property{}

	for _, prop := range properties {
		grpcProp := pb.Property{
			Id:         int32(prop.Id),
			DoorNumber: prop.DoorNumber,
			Address:    prop.Address,
			City:       prop.City,
			Country:    prop.Country,
		}

		grpcProps = append(grpcProps, &grpcProp)
	}

	fmt.Println(properties)

	for _, prop := range grpcProps {
		fmt.Println(fmt.Sprintf("%+v\n", prop))

		if err := stream.Send(prop); err != nil {
			return err
		}
	}

	return nil
}

func (a *Api) BookPropertyById(ctx context.Context, booking *pb.Booking) (*pb.UserPropertyBooking, error) {
	log.Println(
		fmt.Sprintf(
			"Recieved booking request User Id: %d, PropertyId: %d",
			booking.UserId,
			booking.PropertyId,
		),
	)

	user, newBooking, property, err := a.dbService.BookPropertyById(
		int(booking.PropertyId),
		int(booking.UserId),
		booking.StartDate.AsTime(),
		booking.EndDate.AsTime(),
	)

	var (
		pbUser         *pb.User
		pbBooking      *pb.Booking
		pbProperty     *pb.Property
		bookingDetails *pb.UserPropertyBooking
	)

	if err != nil {
		log.Println(fmt.Sprintf("Error saving to database: %v", err))
		return bookingDetails, err
	}

	log.Println(
		fmt.Sprintf(
			"Created Booking User Id: %d, PropertyId: %d booking: %v",
			booking.UserId,
			booking.PropertyId,
			newBooking,
		),
	)

	pbUser, err = user.ConvertUserToMsg()
	if err != nil {
		return bookingDetails, err
	}

	pbBooking, err = newBooking.ConvertBookingToMsg()
	if err != nil {
		return bookingDetails, err
	}

	pbProperty, err = property.ConvertPropertyToMsg()
	if err != nil {
		return bookingDetails, err
	}

	bookingDetails = &pb.UserPropertyBooking{
		User:     pbUser,
		Property: pbProperty,
		Booking:  pbBooking,
	}
	log.Println(fmt.Sprintf("%+v\n", bookingDetails))

	return bookingDetails, nil
}
