package db

import (
	"context"
	"example/GoPractice/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddToMongoDBDatabase(userToInsert *models.User) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Access a specific database and collection
	database := client.Database("mydb")
	collection := database.Collection("mycollection")

	_, err = collection.InsertOne(context.TODO(), userToInsert)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User inserted successfully!")

}
