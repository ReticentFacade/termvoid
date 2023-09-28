package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/ReticentFacade/termvoid/pkg/proto"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

type server struct {
	// UnimplementedFileServiceServer type needs to be embedded
	pb.UnimplementedFileServiceServer
}

func ServerStartup() {
	// Load env variables:
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	// Retrieve env variable:
	PORT := os.Getenv("GRPC_PORT")

	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", ":"+PORT) // ":" before the PORT creates a valid network address
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// Registering service methods from upload.go & download.go
	pb.RegisterFileServiceServer(grpcServer, &server{})
	fmt.Println("Registered service methods successfully!")
	// TODO: Add more from info.go and history.go later

	log.Printf("Server listening on port %s...\n", PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	fmt.Println("Server started successfully...")
}
