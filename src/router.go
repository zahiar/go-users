package main

import (
	"github.com/gorilla/mux"

	"go-users/src/routes"
)

func getRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/status", routes.HandleStatus).
		Methods("GET")
	router.HandleFunc("/user/{id}", routes.HandleUser).
		Methods("GET")

	return router
}
