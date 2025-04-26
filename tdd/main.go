package main

import (
	"log"
	"net/http"

	"Go-Concepts/tdd/dependencies"
)

func main() {
	log.Printf("starting server at localhost:5001...")

	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(dependencies.MyGreeterHandler)))
}
