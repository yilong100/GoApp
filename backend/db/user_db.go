package db

import (
	"bufio"
	"database/sql"
	"example/GoPractice/models"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToPostgresCloudServerAndDB() (*sql.DB, error) {

	filePath := "./db/database-ip-address.txt"
	databaseIP := ""
	// Open the file.
	f, _ := os.Open(filePath)
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		databaseIP = scanner.Text()
		fmt.Print(databaseIP)
	}

	host := databaseIP
	port := 5432
	user := "goApp"
	password := "goApp1234"
	dbname := "users"

	// Create a connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Printf("ERROR 1: %s\n", err) //Exit the program if there's an error
	}

	// // Ping the database to ensure the connection is working
	// err = db.Ping()
	// if err != nil {
	// 	fmt.Printf("ERROR 2: %s\n", err)

	// 	return db, err
	// }

	// Create a table named "users."
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(50) NOT NULL,
		name VARCHAR (50) NOT NULL,
		age INTEGER NOT NULL,
		dreamplacetolive VARCHAR(100) NOT NULL
	);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Printf("ERROR 3: %s\n", err)

		return db, err
	}

	fmt.Println("Table 'users' created successfully.")

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

func CloseConnection(db *sql.DB) bool {
	db.Close()

	return true
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
