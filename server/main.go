package main

import (
	"fmt"
	"log"
	"net"
	"github.com/laughtt/playingRPC"
	"google.golang.org/grpc"
)

func main() {
	//listener TCP PORT 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen")
	}
	//server instance
	s := api.Server{}

	//create a grpc  server object
	grpcServer := grpc.NewServer()

	//atach the ping service to the server

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
