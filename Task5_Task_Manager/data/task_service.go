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

func changeIdToObjectId(id string) (primitive.ObjectID, error) {
	// change id to ObjectId
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id format")
	}

	return objectID, nil
}

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

// GetTasks returns a slice of all tasks in the TaskManager.
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

func (tm *TaskManager) UpdateTask(newTask Task, id string) (TaskResponse, error) {
	objectID, err := changeIdToObjectId(id)
	if err != nil {
		return TaskResponse{}, err
	}

	// Populate update fileds
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
