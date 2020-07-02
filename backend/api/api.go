package api

import (
	"context"
	"log"

	pb "bookings"
)

type Api struct {
	pb.UnimplementedBookingsApiServer
	properties []*pb.Property
}

func New() *Api {
	api := &Api{}
	return api
}

func (a *Api) GetProperties(c context.Context, p *pb.Property, stream pb.BookingsApi_GetAllPropertiesServer) error {

	log.Println(p)
	properties := [
		&pb.Property{},
	]
	for _, property := range properties {
		if err := stream.Send(); err != nil {
			return err
		}
	}

	return nil
}
