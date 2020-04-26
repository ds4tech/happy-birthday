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

var db *mongo.Database

const DBNAME = "people"
const CONNECTIONSTRING = "mongodb://localhost:27017"
const COLLECTIONNAME = "hellopeople"

func init() {
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	//client, err := mongo.Connect(ctx, options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db = client.Database(DBNAME)
}

func SaveCollection(man HelloMan) {
	col := db.Collection(COLLECTIONNAME)

	//ash := HelloMan{1, "Ash", "Pallet Town"}
	//john := HelloMan{2, "John", "Town"}
	man.Id = 0 //filed not used
	insertResult, err := col.InsertOne(context.TODO(), man)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

func DeleteCollection(person HelloMan) {
	_, err := db.Collection(COLLECTIONNAME).DeleteOne(context.Background(), person, nil)
	// _, err := db.Collection(COLLECTIONNAME).DeleteOne(context.Background(), bson.D{{"Id": 1, "Name": Olga, "dateOfBirth": "1988 - 04 - 12"}}, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func UpdateCollection(newMan HelloMan) {
	oldMan := FindMan(newMan.Name)
	fmt.Println("Updating: ", oldMan.Name)

	col := db.Collection(COLLECTIONNAME)

	_, err := col.UpdateOne(context.TODO(), oldMan, bson.D{{"$set", newMan}})
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
	col := db.Collection(COLLECTIONNAME)

	cursor, err := col.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {
			err := cursor.Decode(&man)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				if man.Name == name {
					return man
				}
			}
		}
	}
	defer cursor.Close(ctx)

	return HelloMan{}
}
