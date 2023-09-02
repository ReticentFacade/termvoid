package main 

import (
	"net/http"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// ====================
		// uploading logic here
		// ====================
	})

	// Start the server
	port := 8080
	fmt.Println("Server started at port: ", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error while starting server: ", err)
		return
	}
}