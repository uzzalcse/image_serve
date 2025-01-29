package main

import (
	"log"
	"net/http"
	_"path/filepath"
)

func main() {
	// Use a relative path to the "images" folder
	fileServer := http.FileServer(http.Dir("../images"))
	http.Handle("/images/", http.StripPrefix("/images/", fileServer))

	log.Printf("Starting server on :8082")
	// Start the HTTP server on port 8082
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}
