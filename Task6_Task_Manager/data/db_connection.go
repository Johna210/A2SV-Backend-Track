package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client         *mongo.Client
	UserCollection *mongo.Collection
	TaskCollection *mongo.Collection
)

func Init() {
	// Initialize the database connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Assign the collection to the UserCollection variable
	Client = client
	UserCollection = client.Database("taskManager").Collection("users")
	TaskCollection = client.Database("taskManager").Collection("tasks")

	log.Printf("userCollection %v", UserCollection)
	log.Printf("taskCollection %v", TaskCollection)

	log.Println("Connected to MongoDB!")

}

// CloseDB closes the database connection.
// It disconnects the client from the database and logs a message when the database is closed.
// Returns an error if there was a problem disconnecting from the database.
func CloseDB() error {
	if err := Client.Disconnect(context.Background()); err != nil {
		return err
	}

	log.Print("Database Closed")

	return nil
}
