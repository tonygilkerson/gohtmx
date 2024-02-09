package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tonygilkerson/gohtmx/internal/gohtmx"
)

func main() {

	// Log to the console with date, time and filename prepended
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//
	// Get environment Variables
	//
	contextFile, exists := os.LookupEnv("WEB_ROOT")
	if exists {
		log.Printf("Using environment variable WEB_ROOT: %v", contextFile)
	} else {
		contextFile = "/etc/gohtmx"
		log.Printf("WEB_ROOT environment variable not set, using default value: %v", contextFile)
	}

	//
	// Create Server
	//
	mux := http.NewServeMux()
  svr := &http.Server{Addr: ":8080", Handler: mux,}

	
	//
	// Define routes
	//
	mux.HandleFunc("/", gothmx.HomeHandler)

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		_,_ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_,_ = w.Write([]byte("ok"))
	})
	
	log.Printf("Listening on: %v", svr.Addr)
	log.Fatal(svr.ListenAndServe())
}
