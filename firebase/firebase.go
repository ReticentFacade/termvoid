package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
)

// Initialize Firebase client, set up configurations, and return it
func InitFirebase() (*firebase.App, error) {
	// Firebase initialization code here
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatal("Error initializing firebase app: ", err)
	}

	return app, err
}
