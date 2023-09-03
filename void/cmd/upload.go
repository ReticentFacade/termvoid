package cmd

import (
	"os"
	"fmt"
	"void/utils"
	"net/http"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload [FILE]",
	Long: `Upload a file to the server.
	Usage:
	termvoid upload [FILE]`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		// Check if file exists
		fileExists, err := utils.IfFileExists(filePath)
		if err != nil {
			fmt.Println("Error while checking file existence: ", err)
			return
		}
		if !fileExists {
			fmt.Println("File does not exist at: ", filePath)
			return
		}

		// ====================
		// uploading logic here
		url := "http://localhost:8080/upload"
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Error while opening file: ", err)
			return
		}
		defer file.Close()

		// Create a multipart/form-data req:
		client := &http.Client{}
		req, err := http.NewRequest("POST", url, file)
		if err != nil {
			fmt.Println("Error while creating request: ", err)
		}
		fmt.Println("Request: ", req)
		req.Header.Set("Content-Type", "multipart/form-data")

		// Send the req: 
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error while sending request: ", err)
			return
		}
		defer resp.Body.Close()

		// Check the resp: 
		if resp.StatusCode == http.StatusOK {
			fmt.Println("File uploaded successfully! 😁 \nStatus: ", resp.StatusCode)
		} else {
			fmt.Println("Error while uploading file! 😢 \nStatus: ", resp.StatusCode)
		}
		// ====================

		fmt.Println("File uploaded successfully!")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
