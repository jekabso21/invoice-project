package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongo establishes a connection to the MongoDB database
func ConnectMongo() {
	uri := os.Getenv("MONGO_URI")

	// Set client options with the provided URI
	clientOptions := options.Client().ApplyURI(uri)

	// Create a new MongoDB client
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v\n", err)
	}

	// Set a timeout for the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v\n", err)
	}

	// Verify the connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v\n", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB successfully!")
}