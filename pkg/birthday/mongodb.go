package birthday

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

	insertResult, err := connection.collection.InsertOne(context.TODO(), man)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func UpdateCollection(oldMan HelloMan, newMan HelloMan) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	connection := Connection{
		collection: client.Database("people").Collection(collectionName),
	}

	findOne := connection.collection.FindOne(context.TODO(), oldMan)

	upsertResult, err := connection.collection.UpdateOne(context.TODO(), findOne, newMan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated a single document: ", upsertResult.UpsertedID)
}

//func FindCollection(name string, birthday string) *mongo.SingleResult {
func FindCollection(man HelloMan) bool {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	col := client.Database("people").Collection(collectionName)
	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {
			//var result bson.M
			var result HelloMan
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				//fmt.Println("\nresult type:", reflect.TypeOf(result))
				if result.Name == man.Name {
					fmt.Println("found:", result.Name)
					man.DateOfBirth = result.DateOfBirth
					return true
				}
			}
		}
	}

	//err = col.FindOne(context.TODO(), bson.D{}).Decode(&result)
	//fmt.Printf("findOne: %v\n", result)
	return false
}
