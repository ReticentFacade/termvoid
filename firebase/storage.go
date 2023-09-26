package firebase

import (
	"context"
	"fmt"
	"log"

	"os"

	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	// "firebase.google.com/go/v4/auth"
	// "google.golang.org/api/option"
)

func InitializeFirebase() (*firebase.App, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return nil, err
	}

	// Fetch the STORAGE_BUCKET_NAME from the environment variables
	storageBucketName := os.Getenv("STORAGE_BUCKET_NAME")

	// Setting up storage configurations:
	config := &firebase.Config{
		StorageBucket: storageBucketName,
	}
	fmt.Println("termvoid/firebase/storage.go: Storage bucket name: ", config.StorageBucket)

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
		return nil, err
	}

	return app, nil
}
