package main

import (
	"fmt"
	"github.com/laughtt/playingRPC/api"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//listener TCP PORT 7777
	lis, err := net.listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("Failed to listen")
	}
	//server instance
	s := api.Server{}

	//create a grpc  server object
	grpcServer := grpc.NewServer()

	//atach the ping service to the server
	api.RegisterPingServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
