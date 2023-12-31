package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"termvoid/config"

	pb "termvoid/proto"

	// "github.com/joho/godotenv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	// UnimplementedFileServiceServer type needs to be embedded
	pb.UnimplementedFileServiceServer
}

// Exported function to create an instance of the Server [`NewServer()` = a constructor for the Server type]:
func NewServer() *server {
	return &server{}
}

func UploadFile(stream pb.FileService_UploadFileServer) (err error) {
	fmt.Println("UploadFile() invoked...")
	for {
		fileChunk, err := stream.Recv()
		fmt.Println("Received file chunk", fileChunk)
		if err != nil {
			if err == io.EOF {
				// Return nil to indicate successful completion
				return nil
			}
			return err
		}

		// Implement upload logic
		fmt.Printf("Received %d bytes of data\n", len(fileChunk.Data))
	}
}

func ServerStartup() {
	// Load env variables:
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	PORT := config.GetGRPCPort()
	fmt.Println("PORT---->", PORT)

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

	s := grpc.NewServer()
	// Register reflection service on gRPC server.
	reflection.Register(s)
	pb.RegisterFileServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// fmt.Println("Server started successfully...")
}

// Other methods for GetFileInfo, DownloadFile, and GetFileHistory...
