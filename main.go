package main

import (
	"log"
	"net"

	endpoint "example.com/go-inventory-grpc/internal/endpoint"
	//repository "example.com/go-inventory-grpc/internal/repository"
	config "example.com/go-inventory-grpc/config"

	"google.golang.org/grpc"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "inventory"
)

func main() {

	//ctx := context.Background()

	//initiate Ent Client
	log.Printf("Inventory Service Initilizing Database Connection.......")
	client, err := config.NewEntClient()
	if err != nil {
		log.Printf("err : %s", err)
	}
	defer client.Close()

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)

	// entClient, err := repository.NewPostgres(ctx, psqlInfo)
	// if err != nil {
	// 	log.Println("failed to setup ent client")
	// }
	// tx, err := entClient.Tx(ctx)

	// repository.WithTx(ctx, tx)

	//set the client to the variable defined in package config
	//this will enable the client intance to be accessed anywhere through the accessor which is a function
	//named GetClient
	config.SetClient(client)
	log.Printf("Inventory Service Successfully made Database Connection.......")

	log.Printf("Inventory Service Starting.......")
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen on port 8000: %v", err)
	}

	e := endpoint.Server{}

	grpcServer := grpc.NewServer()

	endpoint.RegisterInventoryServiceServer(grpcServer, &e)

	// msg, err := endpoint.Register(context.Background(), "hi")
	// if err != nil {
	// 	log.Println("failed to call register")
	// }

	//fmt.Println(msg)

	log.Printf("Inventory Service Starting....at port...9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc server over port 9000: %v", err)
	}
	log.Printf("Inventory Service Started Succesfully....at port...9000")
}

// func constructInterceptors() []grpc.UnaryServerInterceptor {

// 	trasactionMiddleware := e
// }

// func main() {
// 	log.Println("Starting listening on port 8080")
// 	port := ":8080"

// 	lis, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	log.Printf("Listening on %s", port)
// 	srv := server.NewGRPCServer()

// 	if err := srv.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }
