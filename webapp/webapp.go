package main

import (
	"fmt"
	"net/http"
)

// homePage function which will take in
// two arguments: http.ResponseWriter and
// a pointer to http.Request
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World...")
}

// setupRoutes function maps incoming requests to
// their intended HTTP handler functions
func setupRoutes() {
	http.HandleFunc("/", homePage)
}

// This will print out a string indicating that the application
// has started. It will then call the setupRoutes function before
// listening and serving the Go application on port 3000.
func main() {
	fmt.Println("Go Web App Started on Port 3000")
	setupRoutes()
	http.ListenAndServe(":3000", nil)
}
