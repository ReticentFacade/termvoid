package localhost

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
	"void/utils"
)

func UploadingFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/upload" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Println("server.go: Uploading file...")
	// ====================
	// uploading logic here
	// Extracting metadata

	// Extracting name from the req header:
	fileName := filepath.Base(r.Header.Get("Content-Disposition"))
	fmt.Println("server.go: Name:", fileName)

	// Assign ID, Downloads, and Expiration Date to the file:
	fileID, err := utils.GenerateID()
	if err != nil {
		fmt.Println("Error while generating ID:", err)
		return
	}
	fmt.Println("server.go: ID:", fileID)
	downloadsAllowed := 10
	fmt.Println("server.go: Downloads allowed:", downloadsAllowed)
	// expirationDate := utils.SetExpirationDate()
	expirationDate := time.Now().Add(24 * time.Hour)
	fmt.Println("server.go: Expiration date: ", expirationDate)
	// ====================
}

func DownloadingFunc(w http.ResponseWriter, r *http.Request) {
	// Check if the req is valid for download:
	if r.URL.Path != "/download" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Println("server.go: Downloading file...")

	// ====================
	// downloading logic here

	// 1. Extract fileName from req or use a query param:
	// fileName := r.URL.Query().Get("file")

	// 2. Retrieve fileContent or metaData (according to the req)

	// 3. Use http.ResponseWriter to send the fileContent to the client
	// 4. Use http.ServeFile or write the fileContent directly to w (http.ResponseWriter)

	// 5. Update downloading counts in the database
	// 6. Update the file's lastDownloadedAt in the database
	// ====================
}
