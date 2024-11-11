package main

import (
	"CodeLineCounter/LineCounter"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// main function sets up the HTTP server
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/count", LineCounter.CountHandler).Methods("GET")

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
