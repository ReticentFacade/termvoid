package cmd

import (
	"golang.org/x/net/context"
	// "context"
	"fmt"

	pb "termvoid/proto"

	"github.com/spf13/cobra"
)

// // infoCmd represents the info command
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
	// Implement logic to fetch file information based on the request
	// You can return a FileMetadata message with the file details
	// or handle it as needed
	return nil, nil
}

func init() {
	rootCmd.AddCommand(infoCmd)
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
