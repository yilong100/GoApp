package server

import (
	"github.com/gorilla/mux"
)

// starts the server
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	routes(router)
	return router
}
