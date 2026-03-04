package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/status", statusHandler)
	mux.HandleFunc("GET /ip", handlerIp)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
