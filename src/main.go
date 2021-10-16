package main

import (
	"log"
	"net/http"
)

func main() {
	router := getRouter()

	log.Println("Listening for connections on 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
