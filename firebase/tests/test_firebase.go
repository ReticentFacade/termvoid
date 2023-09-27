package main

import (
	"fmt"
	"log"

	// "termvoid/firebase"
	"github.com/ReticentFacade/termvoid/firebase"
	// firebase "firebase.google.com/go/v4"
)

func main() {
	fmt.Println("firebaseCheck.go: Firebase check")
	app, err := firebase.InitializeFirebase()
	if err != nil {
		log.Fatal("Error initializing firebase app: ", err)
	}

	fmt.Println("Firebase app initialized successfully!")
	fmt.Println("Firebase app initialized: ", app)
}
