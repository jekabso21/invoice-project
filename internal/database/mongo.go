package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectMongo() {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		log.Fatalf("MongoDB URI not found in environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v\n", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v\n", err)
	}

	MongoClient = client
	log.Println("Connected to MongoDB successfully!")
}

func InsertOne(database, collection string, document interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := MongoClient.Database(database).Collection(collection)
	result, err := coll.InsertOne(ctx, document)
	return result, err
}

func FindOne(database, collection string, filter interface{}) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := MongoClient.Database(database).Collection(collection)
	var result bson.M
	err := coll.FindOne(ctx, filter).Decode(&result)
	return result, err
}

func UpdateOne(database, collection string, filter, update interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := MongoClient.Database(database).Collection(collection)
	result, err := coll.UpdateOne(ctx, filter, bson.M{"$set": update})
	return result, err
}

func DeleteOne(database, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	coll := MongoClient.Database(database).Collection(collection)
	result, err := coll.DeleteOne(ctx, filter)
	return result, err
}
