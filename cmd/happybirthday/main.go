package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ds4tech/happy-birthday/pkg/birthday"
)

var file = birthday.CreateFile("info.log")

func main() {
	fmt.Println("Starting server...")

	router := birthday.NewRouter(file)
	log.Fatal(http.ListenAndServe(":8888", router))

	defer birthday.CloseFile(file)
}
