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
	webRoot, exists := os.LookupEnv("WEB_ROOT")
	if exists {
		log.Printf("Using environment variable WEB_ROOT: %v", webRoot)
	} else {
		webRoot, _ = os.Executable() 
		log.Printf("WEB_ROOT environment variable not set, using default value: %v", webRoot)
	}

	//
	// Create Server
	//
	mux := http.NewServeMux()
  svr := &http.Server{Addr: ":8080", Handler: mux,}

	// Create a new site instance
	site := gothmx.NewSite(webRoot)

	//
	// Define routes
	//
	mux.HandleFunc("/", site.HomeHandler)
	
	mux.HandleFunc("/hi", gothmx.HiHandler)
	mux.HandleFunc("/greeting", gothmx.GreetingHandler)

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		_,_ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_,_ = w.Write([]byte("ok"))
	})
	
	log.Printf("Listening on: %v", svr.Addr)
	log.Fatal(svr.ListenAndServe())
}
