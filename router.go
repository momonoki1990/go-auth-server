package main

import (
	"go-auth-server/auth"
	"go-auth-server/controllers"

	"github.com/gorilla/mux"
)

// Define routing
func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	private := router.NewRoute().Subrouter()

	router.HandleFunc("/public", controllers.PublicHandler)
	router.HandleFunc("/token", controllers.TokenHandler).Methods("POST")

	// Register private handler
	private.HandleFunc("/private", controllers.PrivateHandler)
	private.Use(auth.SampleMiddleWare)

	return router
}
