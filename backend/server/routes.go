package server

import (
	"example/GoPractice/handlers"

	"github.com/gorilla/mux"
)

func routes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.UserById).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
}
