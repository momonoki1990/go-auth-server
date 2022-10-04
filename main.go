package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Start server
	router := initRouter()
	port := ":8080"
	fmt.Printf("server is listening on url: %v", "http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}