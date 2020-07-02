package main

import (
	"fmt"
	"log"
	"net"
	"net/url"

	pb "github.com/SimonTanner/simple-grpc-app/bookings"
	"github.com/jmoiron/sqlx"

	"github.com/SimonTanner/simple-grpc-app/backend/api"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port   = 8090
	dbPort = 5432
	host   = "api"
)

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			"user", url.QueryEscape("password"),
			"db", dbPort, "bookings"))
	if err != nil {
		log.Fatalf("Error connecting to the database %v", err)
	}

	api := api.New(db)

	grpcServer := grpc.NewServer()
	pb.RegisterBookingsApiServer(grpcServer, api)
	err = grpcServer.Serve(lis)

	if err != nil {
		log.Print(err)
	}
}
