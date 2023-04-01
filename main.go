package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	config "example.com/go-inventory-grpc/config"
	staffExternal "example.com/go-inventory-grpc/internal/domain/external/staff"
	staffManager "example.com/go-inventory-grpc/internal/domain/identifiers"
	"example.com/go-inventory-grpc/internal/domain/identifiers/creators"
	staffDomain "example.com/go-inventory-grpc/internal/domain/staff"
	"example.com/go-inventory-grpc/internal/repository"
	staffRepo "example.com/go-inventory-grpc/internal/repository/staff"
	staffClient "example.com/go-inventory-grpc/internal/service/staff"
	server "example.com/go-inventory-grpc/server"
)

func main() {

	ctx := context.Background()

	log.Printf("Inventory Service Initilizing Database Connection.......")
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("failed to load in configuration: ", err)
		os.Exit(1)
	}

	postgresURL := cfg.Database.PostgresURL
	if postgresURL == "" || postgresURL == "null" {
		postgresURL = fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
			cfg.DBConnection.User, cfg.DBConnection.Password, cfg.DBConnection.Host, cfg.DBConnection.Port, cfg.DBConnection.Dbname)
	}

	db, err := repository.NewPostgres(ctx, postgresURL)
	if err != nil {
		fmt.Println("Failed to set up  DB")
	}

	httpClient := http.DefaultClient
	httpClient.Timeout = 100 * time.Second

	// staff client
	staffClient := staffClient.New(httpClient, cfg.ExternalAPI.Host)

	staffExternal := staffExternal.New(staffClient)

	staffManager := staffManager.New(creators.Config{
		StaffExternal: staffExternal,
	})

	staffRepo := staffRepo.New(db)

	staffD := staffDomain.New(staffRepo, staffExternal, staffManager)

	serv := server.New(db, cfg, staffD)

	if err := serv.Start(); err != nil {
		fmt.Println("Failed to start service")
	}

	// go func() {
	// 	if err := serv.Start(); err != nil {
	// 		fmt.Println("Failed to start service")
	// 	}
	// }()

}
