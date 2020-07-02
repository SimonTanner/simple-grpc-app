package api

import (
	"fmt"
	"log"

	pb "github.com/SimonTanner/simple-grpc-app/bookings"
)

type Api struct {
	pb.UnimplementedBookingsApiServer
	properties []*pb.Property
}

func New() *Api {
	api := &Api{}
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
