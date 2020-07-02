package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/SimonTanner/simple-grpc-app/bookings"

	"github.com/SimonTanner/simple-grpc-app/backend/api"

	"google.golang.org/grpc"
)

const (
	port = 8090
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	api := api.New()

	grpcServer := grpc.NewServer()
	pb.RegisterBookingsApiServer(grpcServer, api)
	grpcServer.Serve(lis)

	if err != nil {
		log.Print(err)
	}
}
