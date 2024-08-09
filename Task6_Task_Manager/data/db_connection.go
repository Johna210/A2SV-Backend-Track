package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client         *mongo.Client
	TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
}

var DB = &Database{}

// ConnectToMongoDB connects to a MongoDB database using the provided URI, database name, and collection name.
// It returns a new instance of TaskManager and an error if any occurred during the connection process.
func (db *Database) ConnectToMongoDB(uri, dbName string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	TaskCollection := client.Database(dbName).Collection("tasks")
	UserCollection := client.Database(dbName).Collection("users")
	return &Database{
		Client:         client,
		TaskCollection: TaskCollection,
		UserCollection: UserCollection,
	}, nil
}

// CloseDB closes the database connection.
// It disconnects the client from the database and logs a message when the database is closed.
// Returns an error if there was a problem disconnecting from the database.
func (db *Database) CloseDB() error {
	if err := db.Client.Disconnect(context.Background()); err != nil {
		return err
	}

	log.Print("Database Closed")

	return nil
}
