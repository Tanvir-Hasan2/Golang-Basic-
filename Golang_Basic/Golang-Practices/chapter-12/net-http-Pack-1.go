package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Create a new HTTP handler.
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Write a response to the client.
		fmt.Fprintf(w, "Hi! I am Anik Adnan")
	}

	// Register the handler with the HTTP server.
	http.HandleFunc("/", handler)

	// Start the HTTP server.
	fmt.Println("Starting HTTP server on port 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
