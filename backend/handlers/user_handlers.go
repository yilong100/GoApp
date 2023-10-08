package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yilong100/GoApp/models"
	"github.com/yilong100/GoApp/db"
	"github.com/yilong100/GoApp/utils"
)

// get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Marshal the array to JSON
	jsonResponse, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate that we're sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response writer
	w.Write(jsonResponse)
}

// returns the user by id
func userById(w http.ResponseWriter, r *http.Request) {
	// extracts id url parameter
	params := mux.Vars(r)
	id := params["id"]

	// calling helper function to return user object
	user, err := getUserById(id)
	//if not found throw error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate that we're sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Marshal the user object to json
	jsonResponse, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response to the HTTP response writer
	w.Write(jsonResponse)

}

// helper function to find user with matching id
func getUserById(id string) (*user, error) {
	for index, user := range users {
		if user.ID == id {
			return &users[index], nil
		}
	}

	return nil, errors.New("user not found")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a user object
	var newUser user
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Process the request
	response := ResponseObject{
		Message: "Welcome, " + newUser.Name + "!",
	}

	// Encode the response object to JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	//adding new user to users array
	users = append(users, newUser)
}