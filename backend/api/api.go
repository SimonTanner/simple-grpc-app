package api

import (
	"fmt"
	"log"

	"github.com/SimonTanner/simple-grpc-app/backend/bookings"
	pb "github.com/SimonTanner/simple-grpc-app/bookings"
	"github.com/jmoiron/sqlx"
)

type Api struct {
	pb.UnimplementedBookingsApiServer
	dbService  bookings.Service
	properties []*pb.Property
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
	property := pb.Property{
		Id:         123,
		DoorNumber: "23",
		Address:    "Cadogan Terrace",
		City:       "London",
		Country:    "UK",
	}
	properties := []*pb.Property{&property}

	fmt.Println(properties)

	for _, prop := range properties {
		if err := stream.Send(prop); err != nil {
			return err
		}
	}

	return nil
}
