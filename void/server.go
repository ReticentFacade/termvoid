package main

import (
	"log"
	"net"

	pb "github.com/ReticentFacade/termvoid/pkg/proto"

	"google.golang.org/grpc"
)

const (
	port = "9000"
)

type server struct {
	// UnimplementedFileServiceServer type needs to be embedded
	pb.UnimplementedFileServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":"+port) // ":" before the port creates a valid network address
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// app, err := firebase.NewApp(context.Background(), nil)
	// if err != nil {
	// 	log.Fatal("Error initializing firebase app: ", err)
	// }
	// fmt.Println("Firebase app initialized: ", app)

	grpcServer := grpc.NewServer()
	// Registering service methods from upload.go & download.go
	pb.RegisterFileServiceServer(grpcServer, &server{})
	// TODO: Add more from info.go and history.go later

	log.Printf("Server listening on port %s...\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
