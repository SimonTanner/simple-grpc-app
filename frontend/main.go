package main

import (
	"fmt"
	"log"

	"github.com/SimonTanner/simple-grpc-app/bookings"
	"github.com/SimonTanner/simple-grpc-app/frontend/frontend-app"
	"google.golang.org/grpc"
)

const (
	port        = 8080
	backendPort = 8090
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", backendPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := bookings.NewBookingsApiClient(conn)

	frontend := frontend.New(client)

	err = frontend.Start(fmt.Sprintf(":%d", port))

	if err != nil {
		log.Println(err)
	}
}
