package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Printf("Listening")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("bundled"))))
	
	log.Printf("This line shouldn't be reached")
}