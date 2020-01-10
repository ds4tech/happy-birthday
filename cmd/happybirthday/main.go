package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ds4tech/happy-birthday/pkg/birthday"
)


func main() {
	fmt.Println("Starting server...\n")

	router := birthday.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
