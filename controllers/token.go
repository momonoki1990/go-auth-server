package controllers

import (
	"fmt"
	"net/http"
	"os"
)

// Token handler
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	r.ParseForm()
	var email string = r.Form["email"][0]
	var password string = r.Form["password"][0]
	authEmail := os.Getenv("AUTH_EMAIL")
	authPass := os.Getenv("AUTH_PASSWORD")
	fmt.Printf("email: %s password: %s \n", email, password)

	// Check params
	if email != authEmail || password != authPass {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication failed")
		return
	}	
	
	// Return token
	
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "This is token handler")
}