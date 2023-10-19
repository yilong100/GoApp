package db

import (
	"bufio"
	"database/sql"
	"example/GoPractice/models"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToPostgresCloudServerAndDB() (*sql.DB, error) {

	filePath := "database-ip-address.txt"
	databaseIP := ""
	// Open the file.
	f, _ := os.Open(filePath)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		databaseIP = scanner.Text()
	}
	println(databaseIP)
	host := databaseIP
	port := 5432
	user := "postgres"
	password := "password"
	dbname := "users"

	// Create a connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err) //Exit the program if there's an error
	}

	// Ping the database to ensure the connection is working
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func GetAllData(db *sql.DB, tableName string) ([]models.User, error) {
	// SQL query to select all data from a specified table
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query) // Execute the query
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close the result set when done

	var data []models.User

	// Iterate through the rows of the result set
	for rows.Next() {
		var item models.User
		// Scan the row data into a struct
		if err := rows.Scan(&item.ID, &item.Name, &item.Age, &item.DreamPlaceToLive); err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func InsertUserData(db *sql.DB, userToInsert *models.User) error {

	var id = userToInsert.ID
	var name = userToInsert.Name
	var age = userToInsert.Age
	var dreamplacetolive = userToInsert.DreamPlaceToLive
	// Create the SQL statement with placeholders for the data to be inserted
	stmt, err := db.Prepare("INSERT INTO users (id, name, age, dreamplacetolive) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close() // Close the prepared statement when done

	// Execute the prepared statement with the provided data
	_, err = stmt.Exec(id, name, age, dreamplacetolive)
	return err
}

// mongodb local connection, get, insert data

// func ConnectToMongoDB() (*mongo.Client, error) {
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println("Connected to MongoDB!")
// 	return client, nil
// }

// func GetAllData(client *mongo.Client, databaseName, collectionName string) ([]bson.M, error) {
// 	// Access the specified database and collection
// 	database := client.Database(databaseName)
// 	collection := database.Collection(collectionName)

// 	// Define an empty filter to match all documents
// 	filter := bson.M{}

// 	// Execute the query to find all documents
// 	cursor, err := collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(context.TODO())

// 	var results []bson.M

// 	// Iterate through the results and decode each document into a map
// 	for cursor.Next(context.TODO()) {
// 		var result bson.M
// 		if err := cursor.Decode(&result); err != nil {
// 			return nil, err
// 		}
// 		results = append(results, result)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func InsertUserData(client *mongo.Client, userToInsert *models.User) error {
// 	// Access a specific database and collection
// 	database := client.Database("mydb")
// 	collection := database.Collection("mycollection")

// 	_, err := collection.InsertOne(context.TODO(), userToInsert)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("User inserted successfully!")
// 	return nil
// }
