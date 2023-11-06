package firebase

import (
	"golang.org/x/net/context"
	// "context"
	"config"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	// "github.com/joho/godotenv"
	"google.golang.org/api/option"
)

// // Initializing Firebase app WITHOUT any authentication for now.
// // TODO: Add authentication.
func InitializeFirebase(ctx context.Context) (*firebase.App, error) {
	// Load env variables:
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Retrieve env variable:
	// storageBucketName := os.Getenv("STORAGE_BUCKET_NAME")
	storageBucketName := config.GetStorageBucketName()
	fmt.Println("firebase.go: Storage bucket name: ", storageBucketName)

	// serviceAccountKeyPath := os.Getenv("SERVICE_ACCOUNT_KEY_PATH")
	serviceAccountKeyPath := config.GetServiceAccountKeyPath()
	// Set up storage config:
	config := &firebase.Config{
		StorageBucket: storageBucketName,
	}

	// Creating an option to use the service account key JSON file:
	opt := option.WithCredentialsFile(serviceAccountKeyPath)
	fmt.Println("firebase.go: Option: ", opt)

	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		log.Fatal("Error initializing firebase app: ", err)
		return nil, err
	}

	return app, err
}
