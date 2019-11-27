package main

import (
	"fmt"
	"log"
	"net"
	"gitlab.com/pantomath-io/demo-grpc/api"
	"google.golang.org/grpc"
  )

func main() {
	//listener TCP PORT 7777
	lis ,err := net.listen("tcp", fmt.Sprintf(":%d",7777))
	if err != nil{
		log.Fatalf("Failed to listen")
	}
	s := api.Server{}

	grpcServer := grpc.New

}
