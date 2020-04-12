package birthday

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "hellopeople"

type Connection struct {
	collection *mongo.Collection
}

// var Client *mongo.Client

// func ConnectMongoDB() {
// 	// Set client options
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// 	// Connect to MongoDB
// 	Client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Check the connection
// 	err = Client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Connected to MongoDB!")
// }

func SaveCollection(man HelloMan) {

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)

	connection := Connection{
		collection: client.Database("people").Collection(collectionName),
	}
	//ash := HelloMan{1, "Ash", "Pallet Town"}
	//john := HelloMan{2, "John", "Town"}

	// Insert a single document
	insertResult, err := connection.collection.InsertOne(context.TODO(), man)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}
