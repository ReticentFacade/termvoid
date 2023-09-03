package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download a file from the server",
	Long: `Download a file from the server.
	Usage:
	termvoid download [FILE-URL]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("download called...")
		fileURL := args[0]
		fmt.Println("File URL: ", fileURL)

		err := downloadFile(fileURL)
		if err != nil {
			fmt.Println("Error while downloading file:", err)
			return
		}
		fmt.Println("File downloaded successfully! 😁 ")
	},
}

func downloadFile(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Extract the file name from the URL or provide a custom name
	// For example, you can use path.Base(url) to get the filename from the URL.

	fileName := "./tests/DownloadedFiles/downloaded_file"

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-")
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-")
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(downloadCmd)
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
