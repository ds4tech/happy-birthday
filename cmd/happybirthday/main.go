package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ds4tech/happy-birthday/pkg/birthday"
)

var file = birthday.CreateFile("info.log")

func main() {
	//birthday.ConnectMongoDB()

	fmt.Println("Starting server...")
	router := birthday.NewRouter(file)
	log.Fatal(http.ListenAndServe(":8888", router))

	defer birthday.CloseFile(file)

	// Close the connection once no longer needed
	// err := birthday.Client.Disconnect(context.TODO())
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("Connection to MongoDB closed.")
	// }
}
