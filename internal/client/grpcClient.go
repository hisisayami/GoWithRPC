package main

import (
	"context"
	"log"
	"time"

	endpoint "example.com/go-inventory-grpc/internal/endpoint"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("didn't connect to port 9090: %v", err)
	}

	defer conn.Close()

	c := endpoint.NewInventoryServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.Register(ctx, &endpoint.Message{})
	if err != nil {
		log.Fatalf("didn't connect to port 9000: %v", err)
	}
	log.Println("response :", r)

	createStaff, err := c.CreateStaff(context.Background(), &endpoint.CreateStaffRequest{
		Name:  "sitaram",
		Email: "sitaram@gmail.com",
	})
	if err != nil {
		log.Fatalf("failed to create staff: %v", err)
	}

	log.Println(createStaff.Name, createStaff.Email, createStaff.Id)

}
