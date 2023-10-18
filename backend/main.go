package main

import (
	"example/GoPractice/server"
	"net/http"

	"github.com/rs/cors"
)

// starts the servers
func main() {

	// filePath := "frontend-ip-address.txt"
	// frontendIP := ""
	// // Open the file.
	// f, _ := os.Open(filePath)
	// // Create a new Scanner for the file.
	// scanner := bufio.NewScanner(f)
	// // Loop over all lines in the file and print them.
	// for scanner.Scan() {
	// 	frontendIP = scanner.Text()
	// }

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
