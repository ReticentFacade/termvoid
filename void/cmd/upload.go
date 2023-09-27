package cmd

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/ReticentFacade/termvoid/pkg/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload [FILE]",
	Long: `Upload a file to the server.
	Usage:
	void upload [FILE]`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize gRPC connection:

		// Set up a connection to the gRPC server:
		conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
		if err != nil {
			log.Fatal("Failed to dial server: ", err)
		}
		defer conn.Close()

		// Create grpc client:
		client := pb.NewFileServiceClient(conn)

		// Specify path to file that's to be uploaded:
		localFilePath := args[0]

		// Specify destinationPath in Firebase Storage:
		destinationPath := "voidFiles/"

		// Open the local file for reading:
		file, err := os.Open(localFilePath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}
		defer file.Close()

		// ============================================================

		// Create context with a timeout:
		ctx := context.Background()

		// Create gRPC stream for uploading the file:
		stream, err := client.UploadFile(ctx)
		if err != nil {
			log.Fatal("Error creating stream: ", err)
		}
		defer stream.CloseSend()

		// Send the fileName to the server:
		if err := stream.Send(&pb.FileContent{Data: []byte(destinationPath)}); err != nil {
			log.Fatal("Error sending fileName to server: ", err)
		}

		// Create buffer for reading the file in chunks:
		buffer := make([]byte, 1024)

		// Read & send the file chunks to the server:
		for {
			n, err := file.Read(buffer)
			if err != nil {
				log.Fatal("Error reading file: ", err)
			}
			if n == 0 {
				break
			}

			// Send the chunk to the server:
			if err := stream.Send(&pb.FileContent{Data: buffer[:n]}); err != nil {
				log.Fatal("Error sending chunk to server: ", err)
			}
		}

		// ============================================================
		// Receive response from server:
		response, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatal("Error receiving response from server: ", err)
		}

		fmt.Println("Upload status:", response.Status)
	},
}

// FOR PROVIDING FILE_URLS, REFER TO: https://firebase.google.com/docs/storage/admin/start#shareable_urls

func UploadFile(stream pb.FileService_UploadFileServer) error {
	for {
		fileChunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return stream.SendAndClose(&pb.FileResponse{})
			}
			return err
		}

		// Implement upload logic
		fmt.Printf("Received %d bytes of data\n", len(fileChunk.Data))
	}
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().BoolP("copy", "c", false, "Upload a copy of the file (& not the original)")
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
