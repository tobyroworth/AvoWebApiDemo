package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Printf("Listening")
	log.Fatal(http.ListenAndServe("127.0.0.1:80", http.FileServer(http.Dir("bundled"))))
	
	log.Printf("This line shouldn't be reached")
}