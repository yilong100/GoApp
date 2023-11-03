package server

import (
	"example/GoPractice/handlers"

	"github.com/gorilla/mux"
)

// sets the endpoints and creates the routes for each function - similar to controller layer in spring
func routes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.UserById).Methods("GET")
	router.HandleFunc("/createUser", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/deleteUser/{id:[0-9]+}", handlers.DeleteUser).Methods("GET")
}
