package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Johna210/A2SV-Backend-Track/Track5_Task_Manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Task = models.Task
type TaskResponse = models.TaskResponse

type TaskManager struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// changeIdToObjectId converts a string representation of an ObjectID to a primitive.ObjectID.
// It takes a string `id` as input and returns a primitive.ObjectID and an error.
// If the conversion is successful, it returns the converted ObjectID and a nil error.
// If the conversion fails, it returns an empty ObjectID and an error indicating an invalid id format.
func changeIdToObjectId(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id format")
	}

	return objectID, nil
}

// ConnectToMongoDB connects to a MongoDB database using the provided URI, database name, and collection name.
// It returns a new instance of TaskManager and an error if any occurred during the connection process.
func (tm *TaskManager) ConnectToMongoDB(uri, dbName, collectionName string) (*TaskManager, error) {
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

	collection := client.Database(dbName).Collection(collectionName)
	return &TaskManager{
		Client:     client,
		Collection: collection,
	}, nil
}

// GetTasks retrieves all tasks from the task manager.
// It returns a slice of TaskResponse and an error, if any.
func (tm *TaskManager) GetTasks() ([]TaskResponse, error) {

	cursor, err := tm.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var tasks []TaskResponse
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to the task manager.
// It takes a newTask parameter of type Task, which represents the task to be added.
// It returns a TaskResponse and an error.
// The TaskResponse contains the details of the added task.
// If an error occurs during the insertion or decoding process, an empty TaskResponse and the error are returned.
func (tm *TaskManager) AddTask(newTask Task) (TaskResponse, error) {
	response, err := tm.Collection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return TaskResponse{}, err
	}

	newTaskID := response.InsertedID
	res := tm.Collection.FindOne(context.TODO(), bson.M{"_id": newTaskID})

	var taskResponse TaskResponse
	if err := res.Decode(&taskResponse); err != nil {
		return TaskResponse{}, err
	}

	return taskResponse, nil

}

// UpdateTask updates a task with the given ID in the task manager.
// It takes a newTask object containing the updated task details and the ID of the task to be updated.
// It returns a TaskResponse object containing the updated task and an error if any.
func (tm *TaskManager) UpdateTask(newTask Task, id string) (TaskResponse, error) {
	objectID, err := changeIdToObjectId(id)
	if err != nil {
		return TaskResponse{}, err
	}

	// Populate update fields
	updateFields := make(bson.M)
	if newTask.Title != "" {
		updateFields["title"] = newTask.Title
	}

	if newTask.Status != "" {
		updateFields["status"] = newTask.Status
	}

	if newTask.Description != "" {
		updateFields["description"] = newTask.Description
	}

	result, err := tm.Collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return TaskResponse{}, errors.New("unable to update Task")
	}

	fmt.Println(result)

	if result.ModifiedCount == 0 {
		return TaskResponse{}, errors.New("no task found with the given ID")
	}

	var updatedTask TaskResponse
	err = tm.Collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&updatedTask)
	if err != nil {
		return TaskResponse{}, err
	}

	return updatedTask, nil
}

// GetTask retrieves a task from the task manager by its ID.
// It takes an ID string as input and returns a TaskResponse and an error.
// If the task is found, the TaskResponse is populated with the task details.
// If the task is not found, it returns an empty TaskResponse and an error indicating that no task was found with the given ID.
func (tm *TaskManager) GetTask(id string) (TaskResponse, error) {
	objectID, err := changeIdToObjectId(id)
	if err != nil {
		return TaskResponse{}, err
	}

	var task TaskResponse

	result := tm.Collection.FindOne(context.TODO(), bson.M{"_id": objectID})

	if err := result.Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return TaskResponse{}, errors.New("no task found with the given ID")
		}
	}

	return task, nil
}

// RemoveTask removes a task from the task manager based on the given ID.
// It takes the ID of the task as a parameter and returns an error if any.
// If the task with the given ID is not found, it returns an error with the message "no task found with the given ID".
// If the task is successfully deleted, it returns nil.
func (tm *TaskManager) RemoveTask(id string) error {
	objectID, err := changeIdToObjectId(id)
	if err != nil {
		return err
	}

	response, err := tm.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("no task found with the given ID")
		}

		return errors.New("unable to delete Task")
	}

	if response.DeletedCount == 0 {
		return errors.New("unable to delete Task")
	}

	return nil
}
