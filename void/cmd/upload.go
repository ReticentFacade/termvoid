/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"fmt"
	"void/utils"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload [FILE]",
	Long: `Upload a file to the server.
	Usage:
	termvoid upload [FILE]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called...")
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
		// ====================

		fmt.Println("File uploaded successfully!")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
