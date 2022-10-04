package controllers

import (
	"fmt"
	"net/http"
)

// Sample handler
func SampleHandler (w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "This is sample handler")
}