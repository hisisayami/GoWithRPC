package server

import (
	"fmt"
	"log"
	"net"

	"example.com/go-inventory-grpc/config"
	staffDomain "example.com/go-inventory-grpc/internal/domain/staff"
	"example.com/go-inventory-grpc/internal/endpoint"
	"example.com/go-inventory-grpc/internal/repository"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type Server struct {
	db       repository.DB
	cfg      *config.Config
	endpoint endpoint.InventoryServiceServer
}

func New(db repository.DB, cfg *config.Config, staffDomain staffDomain.StaffDomain) *Server {
	return &Server{
		db:  db,
		cfg: cfg,
		endpoint: endpoint.New(endpoint.Config{
			StaffD: staffDomain,
			DB:     db,
		}),
	}
}

func (s *Server) Start() error {
	log.Printf("Inventory Service Starting.......")

	host := fmt.Sprintf("%s:%d", s.cfg.Server.GRPCHost, s.cfg.Server.GRPCPort)

	if host == "" || host == "null" {
		host = ":9090"
	}

	lis, err := net.Listen("tcp", host)
	if err != nil {
		fmt.Printf("failed to listen on port 9000: %v", err)
		return errors.Wrap(err, "failed to open up tcp socket")

	}

	//e := endpoint.server{}

	grpcServer := grpc.NewServer()

	endpoint.RegisterInventoryServiceServer(grpcServer, s.endpoint)

	log.Printf("Inventory Service Starting....at port...9000")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("failed to serve grpc server over port 9000: %v", err)
		return errors.Wrap(err, "failed to open server")

	}
	fmt.Println("Inventory Service Started Succesfully....at port...9090")

	return nil
}
