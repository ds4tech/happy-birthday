package birthday

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "hellopeople"

var Client *mongo.Client

func ConnectMongoDB() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	Client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func SaveCollection(db *mongo.Database) {
	collection := db.Collection(collectionName)
	fmt.Println("tutaj dziala; ", collectionName)

	//ash := HelloMan{1, "Ash", "Pallet Town"}
	john := HelloMan{3, "John", "Town"}

	// Insert a single document
	insertResult, err := collection.InsertOne(context.TODO(), john)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
