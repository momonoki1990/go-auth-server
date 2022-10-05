package controllers

import (
	"fmt"
	"net/http"
)

// Public handler
func PublicHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "This is public handler")
}