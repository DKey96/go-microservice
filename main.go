package main

import (
	"log"
	"net/http"
)

const message = "Hello from the other side"

func main() {
	// Create a new server running on localhost (similar as Node.js express setup)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	// If there was an error when starting a server, it is being returned
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Server failed to start. Reason: %v", err)
	}
}
