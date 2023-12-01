package cmd

import (
	"context"
	"flag"
	"log"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("server_address", "localhost:8080", "The server address in the format of host:port")
)

func ConnectionUtil() (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(context.Background(), *serverAddr, grpc.WithInsecure())
	if err != nil {
		// handle error
		// log.Fatalf("did not connect: %v", err)
		log.Print("did not connect: ")
		return nil, err
	}
	defer conn.Close()

	// client := pb.NewFileServiceClient(conn)
	// fileService, err := client.UploadFile(context.Background(), &pb.FileRequest{})
	// log.Print("fileService: ", fileService)

	return conn, nil
}
