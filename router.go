package main

import (
	"go-auth-server/controllers"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/sample", controllers.SampleHandler)
	return router
}