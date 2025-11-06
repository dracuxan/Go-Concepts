package main

import (
	"log"
	"net/http"

	"Go-Concepts/tdd/dependencies"
)

func main() {
	log.Fatal(http.ListenAndServe(":1313", http.HandlerFunc(dependencies.MyGreeterHandler)))
}
