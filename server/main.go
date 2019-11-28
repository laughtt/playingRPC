package main

import (
	"context"
	"log"
	"net"
	"strings"

	pb "github.com/laughtt/playingRPC/customer"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

//Se usa para implementar  customer.CustomServer
type server struct {
	savedCustomers []*pb.CustomerRequest
}

//Create a new costumer

func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *server) GetCustomers(filter *pb.CustomFilter, stream pb.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != " " {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("FAILED TO LISTEN")
	}

	//CREATE SERVER
	s := grpc.NewServer()

	pb.RegisterCustomerServer(s, &server{})
	s.Serve(lis)
}
