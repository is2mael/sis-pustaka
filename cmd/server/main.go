package main

import (
	"log"	
	"net/http"

	"sis-pustaka/router"
)

func main() {
	// Initialize the router
	r := router.NewRouter()

	// Start the server on port 8080
	log.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}