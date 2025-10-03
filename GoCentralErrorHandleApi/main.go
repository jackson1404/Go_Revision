package main

import (
	"fmt"
	"log"
	"net/http"
)

// Centralized error wrapper
func errorHandler(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			log.Printf("request error: %v", err) // centralized logging
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	}
}

// Business logic: only returns error
func helloHandler(w http.ResponseWriter, r *http.Request) error {
	name := r.URL.Query().Get("name")
	if name == "" {
		return fmt.Errorf("missing 'name' query parameter")
	}
	fmt.Fprintf(w, "Hello %s!", name)
	return nil
}

func main() {
	http.HandleFunc("/hello", errorHandler(helloHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
