package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	// Start server
	router := initRouter()
	port := ":8080"
	fmt.Printf("server is listening on url: %v \n", "http://localhost"+port)
	log.Fatal(http.ListenAndServe(port, router))
}
