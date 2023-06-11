package main

import (
	"log"
	"microservice-in-30-mins/homepage"
	"microservice-in-30-mins/server"
	"net/http"
	"os"
)

var (
	CertFile    = os.Getenv("DK_CERT_FILE")
	KeyFile     = os.Getenv("DK_KEY_FILE")
	ServiceAddr = os.Getenv("DK_SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, "dk-service", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)
	// Create a new server running on localhost (similar as Node.js express setup)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.Server(mux, ServiceAddr)

	// If there was an error when starting a server, it is being returned
	logger.Println("Server starting...")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("Server failed to start. Reason: %v", err)
	}
}
