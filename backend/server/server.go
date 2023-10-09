package server

import (
	"github.com/gorilla/mux"
)

// starts the server and sets the routes
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	routes(router)
	return router
}
