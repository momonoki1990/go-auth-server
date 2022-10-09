package controllers

import (
	"fmt"
	"net/http"
)

// Private handler
func PrivateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Authorization succeeded (successfully accessed to private handler)")
}