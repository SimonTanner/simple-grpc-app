package api

import (
	"context"
	"log"

	pb "github.com/SimonTanner/simple-grpc-app/backend/bookings"
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
	property := pb.Property{}
	properties := []*pb.Property{&property}

	for _, prop := range properties {
		if err := stream.Send(prop); err != nil {
			return err
		}
	}

	return nil
}
