package utils

import (
	"os"
)

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