package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/ReticentFacade/termvoid/pkg/proto"

	"google.golang.org/grpc"
)

const (
	port = "9000"
)

type server struct {
	pb.UnimplementedFileServiceServer // Embed the UnimplementedFileServiceServer type
}

func (s *server) UploadFile(stream pb.FileService_UploadFileServer) error {
	for {
		// Receive chunks of file data from the client
		fileChunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				// Client has finished sending data
				return stream.SendAndClose(&pb.FileResponse{
					// Provide a response indicating the upload was successful
				})
			}
			// Handle other errors
			return err
		}

		// Process the received file chunk (fileChunk.Data)
		// Implement your upload logic here
		fmt.Printf("Received %d bytes of data\n", len(fileChunk.Data))
	}
}

func (s *server) DownloadFile(ctx context.Context, request *pb.FileRequest) (*pb.FileContent, error) {
	// Implement your download logic here based on the request
	// You can return a FileContent message with the file data
	// or handle it as needed
	return nil, nil
}

func (s *server) GetFileInfo(ctx context.Context, request *pb.FileRequest) (*pb.FileMetadata, error) {
	// Implement logic to fetch file information based on the request
	// You can return a FileMetadata message with the file details
	// or handle it as needed
	return nil, nil
}

func (s *server) GetFileHistory(ctx context.Context, request *pb.FileRequest) (*pb.FileHistory, error) {
	// Implement logic to fetch file history based on the request
	// You can return a FileHistory message with the history details
	// or handle it as needed
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+port) // ":" before the port creates a valid network address
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
