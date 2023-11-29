package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	pb "termvoid/proto"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("server_address", "localhost:8080", "The server address in the format of host:port")
)

func main() {
	conn, err := grpc.DialContext(context.Background(), *serverAddr, grpc.WithInsecure())
	if err != nil {
		// handle error
		fmt.Print("Error: ", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	fileService, err := client.UploadFile(context.Background(), &pb.FileRequest{})

	log.Print("fileService: ", fileService)
}
