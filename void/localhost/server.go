package localhost 

import (
	"fmt"
	"net/http"
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
    fmt.Fprintf(w, "Hello World!")
    fmt.Println("Uploading file...")
    // ====================
    // uploading logic here
    // ====================
}