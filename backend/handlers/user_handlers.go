// Contains functions and business logic of application - similar to service layer in spring

package handlers

import (
	"encoding/json"
	"errors"
	"example/GoPractice/models"
	"net/http"

	"example/GoPractice/db"
	"log"

	"github.com/gorilla/mux"
)

// get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Establish a MongoDB connection
	postgresdb, err := db.ConnectToPostgresCloudServerAndDB()
	if err != nil {
		http.Error(w, "Failed to connect to Postgres", http.StatusBadGateway)
		db.CloseConnection(postgresdb)
		return
	}

	// Get all data using the established connection
	allData, err := db.GetAllData(postgresdb, "users")
	if err != nil {
		http.Error(w, "Failed to retrieve data from database", http.StatusBadGateway)
		log.Println("Error:", err)
		db.CloseConnection(postgresdb)
		return
	}
	// Marshal the array to JSON
	jsonResponse, err := json.Marshal(allData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		db.CloseConnection(postgresdb)
		return
	}

	// Set the Content-Type header to indicate that we're sending JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response to the HTTP response writer
	w.Write(jsonResponse)
	db.CloseConnection(postgresdb)
}

// returns the user by id
func UserById(w http.ResponseWriter, r *http.Request) {
	// extracts id url parameter
	params := mux.Vars(r)
	id := params["id"]

	// calling helper function to return user object
	user, err := GetUserById(id)
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
func GetUserById(id string) (*models.User, error) {
	for index, user := range models.Users {
		if user.ID == id {
			return &models.Users[index], nil
		}
	}

	return nil, errors.New("user not found")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into a user object
	var newUser models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	// Process the request
	response := models.ResponseObject{
		Message: "Welcome, " + newUser.Name + "!",
	}

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Encode the response object to JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	//adding new user to users array
	models.Users = append(models.Users, newUser)

	// Establish a MongoDB connection
	postgresdb, err := db.ConnectToPostgresCloudServerAndDB()
	if err != nil {
		http.Error(w, "Failed to connect to PostgresDB", http.StatusBadGateway)
		db.CloseConnection(postgresdb)
		return
	}

	// Insert user data using the established connection
	err = db.InsertUserData(postgresdb, &newUser)
	if err != nil {
		http.Error(w, "Failed to append user to database", http.StatusBadGateway)
		log.Println("Error:", err)
		db.CloseConnection(postgresdb)
		return
	}

	db.CloseConnection(postgresdb)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	// extracts id url parameter
	params := mux.Vars(r)
	id := params["id"]

	// Process the request
	response := models.ResponseObject{
		Message: "Removed User",
	}

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Encode the response object to JSON and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

	// Establish a MongoDB connection
	postgresdb, err := db.ConnectToPostgresCloudServerAndDB()
	if err != nil {
		http.Error(w, "Failed to connect to PostgresDB", http.StatusBadGateway)
		db.CloseConnection(postgresdb)
		return
	}

	// Insert user data using the established connection
	err = db.DeleteUserData(postgresdb, id)
	if err != nil {
		http.Error(w, "Failed to append user to database", http.StatusBadGateway)
		log.Println("Error:", err)
		db.CloseConnection(postgresdb)
		return
	}

	db.CloseConnection(postgresdb)
}
