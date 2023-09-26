package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
)

func UploadFile(app *firebase.App, filename string) error {
	fmt.Println("termvoid/firebase/upload.go: Uploading file: ", filename)
	_, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("Error initializing firebase app: ", err)
	}

	return err
}
