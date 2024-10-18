package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler for endpoint1
func endpoint1(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed endpoint1")
	fmt.Fprintln(w, "Welcome! You've reached endpoint1.")
}

// Handler for endpoint2
func endpoint2(w http.ResponseWriter, r *http.Request) {
	log.Println("Accessed endpoint2")
	fmt.Fprintln(w, "Greetings! You're now at endpoint2.")
}

// Handler for the root path
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Root accessed")
	fmt.Fprintln(w, "Welcome! Use either /endpoint1 or /endpoint2 to explore.")
}

func main() {
	// Defining route handlers
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/endpoint1", endpoint1)
	http.HandleFunc("/endpoint2", endpoint2)

	// Starting the server
	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
