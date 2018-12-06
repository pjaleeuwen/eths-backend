package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := createRouter()

	fmt.Println("running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}