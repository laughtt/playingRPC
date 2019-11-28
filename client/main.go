package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/laughtt/playingRPC/customer"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

//CREATE THE METHOD CREATE COSTUMER OF COSTUMESERVER
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal("CANT CREATE COSTUMER")
	}
	if resp.Success {
		log.Printf("A NEW COSTUMER WAS PRINT WITH ID : %d", resp.Id)
	}
}

func getCostumers(client pb.CustomerClient, filter *pb.CustomFilter) {
	//CALLING THE STREAMING API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatalf("ERROR CANT GET COSTUMER")
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}
func main() {
	// Set up a connection to the gRPC server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(err)
	defer conn.Close()
	// Creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerRequest{
		Id:    104,
		Name:  "jose",
		Email: "sasa",
		Phone: "41",
		Addresses: []*pb.Address{
			&pb.Address{
				Street:            "sas",
				City:              "sf",
				State:             "41",
				Zip:               "sa",
				IsShippingAddress: true,
			},
		},
	}
	createCustomer(client, customer)
	customer = &pb.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.Address{
			&pb.Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.Address{
			&pb.Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)
	// Filter with an empty Keyword
	filter := &pb.CustomFilter{Keyword: ""}
	getCostumers(client, filter)
}
