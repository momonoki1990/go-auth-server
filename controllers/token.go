package controllers

import (
	"fmt"
	"log"
	"net/http"

	"go-auth-server/auth"
)

// Token handler
func TokenHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	r.ParseForm()
	var email string = r.Form["email"][0]
	var password string = r.Form["password"][0]
	fmt.Printf("email: %s password: %s \n", email, password)

	// Check params
	if email != "sample@sample.com" || password != "mypassword" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authentication failed")
		return
	}

	// Generate token
	token, err := auth.GenerateToken()
	if err != nil {
		log.Println("Error at GetToken():", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Internal Server Error")
		return
	}
	
	// Return token as cookie
	cookie := &http.Cookie{
		Name: auth.TokenCookieName,
		Value: token,
		// Path: "",
		Domain: "localhost",
		MaxAge: 60 * 60 * 24, // 1 hour
		// Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,	
	}
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Succesfully token returned")
}