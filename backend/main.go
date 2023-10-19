package main

import (
	"example/GoPractice/server"
	"net/http"

	"github.com/rs/cors"
)

// starts the servers
func main() {

	router := server.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})

	http.Handle("/", c.Handler(router))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
