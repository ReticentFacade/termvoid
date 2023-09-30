package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	pb "github.com/ReticentFacade/termvoid/pkg/proto"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	grpcServer := grpc.NewServer()
	// Register server reflection service:
	reflection.Register(grpcServer)
	// Registering service methods:
	pb.RegisterFileServiceServer(grpcServer, &server{})
	fmt.Println("Registered service methods successfully!")
	// TODO: Add more from info.go and history.go later

	fmt.Println("Starting server...")
	lis, err := net.Listen("tcp", ":"+PORT) // ":" before the PORT creates a valid network address
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening on port %s...\n", PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	fmt.Println("Server started successfully...")
}

func UploadFile(stream pb.FileService_UploadFileServer) error {
	fmt.Println("UploadFile() invoked...")
	for {
		fileChunk, err := stream.Recv()
		fmt.Println("Received file chunk", fileChunk)
		if err != nil {
			if err == io.EOF {
				// return stream.SendAndClose(&pb.FileResponse{})
				return stream.SendAndClose(&pb.FileResponse{Status: "Upload successful"})
			}
			return err
		}

		// Implement upload logic
		fmt.Printf("Received %d bytes of data\n", len(fileChunk.Data))
	}
}

// Other methods for GetFileInfo, DownloadFile, and GetFileHistory...
