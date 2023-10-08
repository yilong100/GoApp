package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/yilong100/GoApp/handlers"
)

func routes(router *mux.Router) {
	router.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id:[0-9]+}", handlers.GetUserByID).Methods("GET")
	router.HandleFunc("/users", handlers.CreateUser).Methods("POST")
}