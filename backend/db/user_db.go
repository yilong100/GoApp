package db

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
