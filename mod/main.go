package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeter() {
	fmt.Println("Hello World!")
	fmt.Println("Listening on port 4000...")
}

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home Page!</h1>"))
}

func main() {
	greeter()

	router := http.NewServeMux()
	router.HandleFunc("GET /", serverHome)

	log.Fatal(http.ListenAndServe(":4000", router))
}
