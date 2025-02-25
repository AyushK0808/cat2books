package config

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}
	FirebaseApp = app
}
