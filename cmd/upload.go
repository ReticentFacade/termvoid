package cmd

// import (
// 	"golang.org/x/net/context"
// 	// "context"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// 	"path/filepath"

	"termvoid/server"

	pb "termvoid/pkg/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// var uploadCmd = &cobra.Command{
// 	Use:   "upload",
// 	Short: "upload [FILE]",
// 	Long: `Upload a file to the server.
// 	Usage:
// 	void upload [FILE]`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// Initialize gRPC connection:

// 		// Set up a connection to the gRPC server:
// 		fmt.Println("Connecting to server...")
// 		// conn, err := grpc.Dial("0.0.0.0:9090", grpc.WithInsecure())
// 		conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
// 		if err != nil {
// 			log.Fatal("Failed to dial server: ", err)
// 		}
// 		defer conn.Close()
// 		fmt.Println("Connected to server successfully!")

// 		// Create grpc client:
// 		// !!!!!!!!!!!!!!!!!!!!!
// 		client := pb.NewFileServiceClient(conn)
// 		// client := pb.NewFileServiceClient(conn)

// 		// Specify path to file that's to be uploaded:
// 		localFilePath := args[0]
// 		fileName := filepath.Base(localFilePath)
// 		fmt.Println("fileName:", fileName)

// 		// Specify destinationPath in Firebase Storage:
// 		destinationPath := "voidFiles/"

// 		// Open the local file for reading:
// 		fmt.Println("Opening file...")
// 		file, err := os.Open(localFilePath)
// 		if err != nil {
// 			log.Fatal("Error opening file: ", err)
// 		}
// 		defer file.Close()
// 		fmt.Println("File opened successfully!")

// 		// ============================================================

// 		// Create context with a timeout:
// 		fmt.Println("Creating context...")
// 		ctx := context.Background()

// 		// ------------------------------------------------------------
// 		// Create an instance of the server type
// 		// srv := server.Server{}
// 		srv := server.NewServer()
// 		fmt.Println("srv:", srv)
// 		// ------------------------------------------------------------

// 		// Create gRPC stream for uploading the file:
// 		stream, err := client.UploadFile(ctx)
// 		// stream, err := server.UploadFile(ctx)
// 		fmt.Println("Stream created successfully...")
// 		fmt.Println("Stream ----> ", stream)
// 		if err != nil {
// 			log.Fatal("Error creating stream: ", err)
// 		}
// 		defer stream.CloseSend()
// 		fmt.Println("Stream created successfully!")

// 		// Send the fileName to the server:
// 		if err := stream.Send(&pb.FileContent{Data: []byte(destinationPath)}); err != nil {
// 			log.Fatal("Error sending fileName to server: ", err)
// 		}

// 		// Create buffer for reading the file in chunks:
// 		buffer := make([]byte, 1024)
// 		fmt.Println("Buffer created successfully!")

// 		// Read & send the file chunks to the server:
// 		for {
// 			n, err := file.Read(buffer)
// 			if err != nil {
// 				if err == io.EOF {
// 					// Reached the EndOfFile, exit the loop
// 					fmt.Println("End of file reached.")
// 					break
// 				} else {
// 					log.Fatal("Error reading file: ", err)
// 				}
// 			}
// 			if n == 0 {
// 				break
// 			}

// 			// Send the chunk to the server:
// 			if err := stream.Send(&pb.FileContent{Data: buffer[:n]}); err != nil {
// 				log.Fatal("Error sending chunk to server: ", err)
// 			}
// 		}
// 		fmt.Println("File sent successfully!")

// 		// ============================================================
// 		// Receive response from server:
// 		fmt.Println("Receiving response from server...")
// 		response, err := stream.CloseAndRecv()
// 		if err != nil {
// 			log.Fatal("Error receiving response from server: ", err)
// 		}

// 		fmt.Println("Upload status:", response.Status)

// 		// Initialize Firebase app and upload the file to Firebase Storage
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(uploadCmd)
// 	uploadCmd.Flags().BoolP("copy", "c", false, "Upload a copy of the file (& not the original)")
// 	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
