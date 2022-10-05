package auth

import (
	"fmt"
	"net/http"
)

// Sample middleware
func SampleMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("Sample Middleware has been called")
		next.ServeHTTP(w, r)
	}) 
}