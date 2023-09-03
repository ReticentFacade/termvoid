package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

// Struct to store FileInfo:
type FileInfo struct {
	ID             string
	Name           string
	Size           int64
	MIME           string
	Downloads      int
	ExpirationDate time.Time
}

// IfFileExists checks if a file exists in the current directory or not
func IfFileExistsInCurrDir(fileName string) (bool, error) {
	// Get curr dir:
	currDir, err := os.Getwd()
	if err != nil {
		return false, err
	}

	// Full path to the file in the currDir:
	fullPath := currDir + string(os.PathSeparator) + fileName

	// Check if file exists:
	_, err = os.Stat(fullPath)
	if err == nil {
		return true, nil // Exists
	}
	if os.IsNotExist(err) {
		return false, nil // Doesn't exist
	}

	return false, err // Error while checking
}

// IfFileExists checks if a file exists in the given path or not
func IfFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // Exists
	}
	if os.IsNotExist(err) {
		return false, nil // Doesn't exist
	}

	return false, err // Error while checking
}

func GenerateID() (string, error) {
	return uuid.New().String(), nil
}

func ExtractFileID(FileURL string) (string, error) {
	return "fileID", nil
}

func GetFileInfo(fileID string) (*FileInfo, error) {
	fmt.Println("utils/file.go: GetFileInfo running: fileID: ", fileID)
	return nil, nil
}
