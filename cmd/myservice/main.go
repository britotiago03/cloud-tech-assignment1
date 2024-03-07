package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest-web-service/internal/api"
)

func main() {
	// Port Assignment
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	// Set up routes for endpoints
	http.HandleFunc("/", api.DefaultHandler)
	http.HandleFunc("/librarystats/v1/bookcount/", api.BookCountHandler)
	http.HandleFunc("/librarystats/v1/readership/{language}", api.ReadershipHandler)
	http.HandleFunc("/librarystats/v1/status/", api.StatusHandler)

	// Server start
	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting myservice: %s\n", err)
	}
}
