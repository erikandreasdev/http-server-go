package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Debugging print message
	fmt.Println("Logs from your program will appear here!")

	// Setting up the HTTP handler
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Write HTTP response body
		w.WriteHeader(http.StatusOK) // Set status code as 200 OK (optional because it's the default)
		_, _ = w.Write([]byte("Hello, world!"))
	})

	// Start the HTTP server
	log.Println("Starting server on port 4221...")
	err := http.ListenAndServe(":4221", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
