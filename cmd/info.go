package cmd

import (
	"log"

	"golang.org/x/net/context"
	// "context"
	"fmt"

	pb "termvoid/proto"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Fetches the metadata of a file",
	Long: `
	Usage:
	void info [FILE-URL]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called")
	},
}

func GetFileInfo(ctx context.Context, request *pb.FileRequest) (*pb.FileMetadata, error) {
	err := ConnectionUtil()
	if err != nil {
		log.Fatal("check connection in termvoid/cmd/client.go: ", err)
	}
	fmt.Println("Connection established...")

	// client := pb.NewFileServiceClient()
	// fmt.Print("`client` from NewFileServiceClient: ", client)

	// fileInfo, err := client.FileMetadata{}

	return nil, nil
}

func init() {
	rootCmd.AddCommand(infoCmd)
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
