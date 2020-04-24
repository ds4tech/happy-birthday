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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	col := client.Database("people").Collection(collectionName)

	//ash := HelloMan{1, "Ash", "Pallet Town"}
	//john := HelloMan{2, "John", "Town"}

	insertResult, err := col.InsertOne(context.TODO(), man)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func DeleteCollection(newMan HelloMan) {
	fmt.Println("Delete smb function.")
}

func UpdateCollection(newMan HelloMan) {
	oldMan := FindMan(newMan.Name)
	fmt.Println("Updating: ", oldMan.Name)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	col := client.Database("people").Collection(collectionName)

	_, err = col.UpdateOne(context.TODO(), oldMan, bson.D{{"$set", newMan}})
	if err != nil {
		panic(err)
	}
}

func FindCollection(man HelloMan) bool {
	if FindMan(man.Name) == (HelloMan{}) {
		return false
	} else {
		return true
	}
}

func FindMan(name string) HelloMan {
	var man HelloMan

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
			err := cursor.Decode(&man)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				//fmt.Println("\nresult type:", reflect.TypeOf(result))
				if man.Name == name {
					return man
				}
			}
		}
	}

	return HelloMan{}
}
