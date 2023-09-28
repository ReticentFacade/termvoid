package firebase

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
)

// // Uploads file to storage using a stream of bytes
// func UploadFileToFirebase(ctx context.Context, app *firebase.App, stream io.Reader, fileName, destinationPath string) error {
// 	// Loading env variables:
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file: ", err)
// 	}

// 	// Get firebase storage client:
// 	client, err := app.Storage(ctx)
// 	if err != nil {
// 		log.Fatal("Error initializing firebase storage client: ", err)
// 		return err
// 	}

// 	// New object to handle uploading:
// 	bucket, err := client.Bucket(os.Getenv("STORAGE_BUCKET_NAME"))
// 	if err != nil {
// 		log.Fatal("Error getting bucket: ", err)
// 		return err
// 	}
// 	obj := bucket.Object(destinationPath)

// 	// Creating a writer for the object:
// 	wc := obj.NewWriter(ctx)

// 	// Copying the stream of bytes to the obj in firebase-storage:
// 	if _, err := io.Copy(wc, stream); err != nil {
// 		log.Fatal("Error while copying stream to firebase storage: ", err)
// 		return err
// 	}

// 	// Closing the writer:
// 	if err := wc.Close(); err != nil {
// 		log.Fatal("Error while closing writer: ", err)
// 		return err
// 	}

// 	fmt.Printf("File %s uploaded to Firebase Storage as %s\n", fileName, destinationPath)
// 	return nil
// }

// Uploads file to storage using a stream of bytes
func UploadFileToFirebase(ctx context.Context, app *firebase.App, reader io.Reader, fileName, destinationPath string) error {
	// Loading env variables:
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Get firebase storage client:
	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatal("Error initializing firebase storage client: ", err)
		return err
	}

	// New object to handle uploading:
	bucket, err := client.Bucket(os.Getenv("STORAGE_BUCKET_NAME"))
	if err != nil {
		log.Fatal("Error getting bucket: ", err)
		return err
	}
	obj := bucket.Object(destinationPath)

	// Creating a writer for the object:
	wc := obj.NewWriter(ctx)

	// Copying the stream of bytes to the obj in firebase-storage:
	if _, err := io.Copy(wc, reader); err != nil {
		log.Fatal("Error while copying stream to firebase storage: ", err)
		return err
	}

	// Closing the writer:
	if err := wc.Close(); err != nil {
		log.Fatal("Error while closing writer: ", err)
		return err
	}

	fmt.Printf("File %s uploaded to Firebase Storage as %s\n", fileName, destinationPath)
	return nil
}
