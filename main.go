package main

import (
	"log"
	"microservice-in-30-mins/server"
	"net/http"
	"os"
)

const message = "Hello from the other side"

var (
	CertFile    = os.Getenv("DK_CERT_FILE")
	KeyFile     = os.Getenv("DK_KEY_FILE")
	ServiceAddr = os.Getenv("DK_SERVICE_ADDR")
)

func main() {
	// Create a new server running on localhost (similar as Node.js express setup)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := server.Server(mux, ServiceAddr)

	// If there was an error when starting a server, it is being returned
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		log.Fatalf("Server failed to start. Reason: %v", err)
	}
}
