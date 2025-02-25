package config

import (
	"context"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App
var (
	firestoreClient *firestore.Client
	once            sync.Once
)

func InitFirebase() {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}
	FirebaseApp = app
}

func FirebaseFirestore() (*firestore.Client, error) {
	var err error
	if firestoreClient != nil {
		return firestoreClient, nil
	}

	// Use service account credentials from JSON file
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	firestoreClient, err := firestore.NewClient(ctx, os.Getenv("FIREBASE_PROJECT_ID"), opt)
	if err != nil {
		log.Println("Error initializing Firestore:", err)
	}
	return firestoreClient, err
}
