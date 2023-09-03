package main

import (
	"fmt"
	"void/cmd"
	"net/http"
	"void/localhost"
)

func main() {
	// CLI
	cmd.Execute()

	// ===============================
	// Localhost server
	fileServer := http.FileServer(http.Dir("./localhost"))
	http.Handle("/", fileServer)
	http.HandleFunc("/upload", localhost.UploadingFunc)

	// Start the server
	port := 8080
	fmt.Println("main.go: Server started at port:", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error while starting server:", err)
		return
	}
}
