package cmd

import (
	"fmt"
	"time"

	"void/utils"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Fetch the metadata of a file",
	Long: `Fetch information about a file
	Usage:
	void info [FILE-URL]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("info called...")
		fileURL := args[0]

		// TODO: Extract file's ID from the URL
		fileID, err := utils.ExtractFileID(fileURL)
		if err != nil {
			fmt.Println("Error while extracting file ID:", err)
			return
		}
		fmt.Println("File ID:", fileID)

		// TODO: Fetch the file's metadata using fileID
		fileInfo, err := utils.GetFileInfo(fileID)
		if err != nil {
			fmt.Println("Error while fetching file info:", err)
			return
		}
		fmt.Println("File info:", fileInfo)

		// Display the metadata:
		fmt.Println("ID:", fileInfo.ID)
		fmt.Println("Name:", fileInfo.Name)
		fmt.Println("Size:", fileInfo.Size)
		fmt.Println("MIME:", fileInfo.MIME) // Type of the file, e.g., "PDF", ".jpg", ".txt", etc.
		fmt.Println("Downloads:", fileInfo.Downloads)
		fmt.Println("Expiration Date:", fileInfo.ExpirationDate.Format(time.RFC3339))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
