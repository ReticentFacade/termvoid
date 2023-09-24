package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ReticentFacade/termvoid/gRPC/proto"
	"google.golang.org/grpc"
)

const (
	port = "9000"
)

type server struct {
	// Implement gRPC methods here
}

func (s *server) UploadFile(ctx context.Context, request *pb.FileContent) (*pb.FileResponse, error) {
	fmt.Println("UploadFile called...")

	// File upload logic here:
	return &pb.FileResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, &server{}) // Register the service implementation

	log.Printf("Server listening on port %s...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
