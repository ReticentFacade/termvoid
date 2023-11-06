package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config file for .env-variables and context initialization:

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return err
}

func GetStorageBucketName() string {
	return os.Getenv("STORAGE_BUCKET_NAME")
}

func GetServiceAccountKeyPath() string {
	return os.Getenv("SERVICE_ACCOUNT_KEY_PATH")
}

func GetGRPCPort() string {
	return os.Getenv("GRPC_PORT")
}
