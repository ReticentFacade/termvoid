package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {
	fmt.Println("downloadTest.go: Downloading file...")

	// Test HTTP server
	go func() {
		http.Handle("/testfiles/", http.StripPrefix("/testfiles/", http.FileServer(http.Dir("tests"))))
		port := 8080
		fmt.Println("Server started at port:", port)
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			fmt.Println("Error while starting server:", err)
			return
		}
	}()

	// Allow the test server to start before running the test
	time.Sleep(1 * time.Second)
}

func TestDownloadingHandler(t *testing.T) {
	// Create a test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("This is a test file content."))
	}))
	defer server.Close()

	// Using test-server's URL as the file URL for `download` cmd
	// fileURL := "http://localhost:8080/testfiles/testFile.txt"
	fileURL := "https://gist.github.com/ReticentFacade/504a8c38d43efcc8333a9b84c4d5cd7b"
	fmt.Println("File URL: ", fileURL)

	// Perform the file download
	resp, err := http.Get(fileURL)
	if err != nil {
		t.Fatalf("Error while downloading file: %v", err)
	}
	defer resp.Body.Close()

	// Save the content to a file in the 'tests/DownloadedFiles' directory
	filePath := "./DownloadedFiles/testFile.html"
	file, err := os.Create(filePath)
	if err != nil {
		t.Fatalf("Error while creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		t.Fatalf("Error while saving file: %v", err)
	}

	fmt.Println("File downloaded successfully and saved as testFile.txt.")
}
