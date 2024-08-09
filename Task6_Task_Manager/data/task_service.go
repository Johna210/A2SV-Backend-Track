package data

import (
	"context"
	"errors"
	"fmt"

	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task = models.Task

// changeIdToObjectId converts a string representation of an ObjectID to a primitive.ObjectID.
// It takes a string `id` as input and returns a primitive.ObjectID and an error.
// If the conversion is successful, it returns the converted ObjectID and a nil error.
// If the conversion fails, it returns an empty ObjectID and an error indicating an invalid id format.
func ChangeIdToObjectId(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, errors.New("invalid id format")
	}

	return objectID, nil
}

// GetTasks retrieves all tasks from the task manager.
// It returns a slice of Task and an error, if any.
func GetTasks() ([]Task, error) {

	cursor, err := TaskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

// AddTask adds a new task to the task manager.
// It takes a newTask parameter of type Task, which represents the task to be added.
// It returns a Task and an error.
// The Task contains the details of the added task.
// If an error occurs during the insertion or decoding process, an empty Task and the error are returned.
func AddTask(newTask Task) (Task, error) {
	// Set the ID field to a new ObjectID
	newTask.ID = primitive.NewObjectID()

	response, err := TaskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return Task{}, err
	}

	newTaskID := response.InsertedID.(primitive.ObjectID)
	res := TaskCollection.FindOne(context.TODO(), bson.M{"_id": newTaskID})

	var task Task
	if err := res.Decode(&task); err != nil {
		return Task{}, err
	}

	return task, nil

}

// UpdateTask updates a task with the given ID in the task manager.
// It takes a newTask object containing the updated task details and the ID of the task to be updated.
// It returns a Task object containing the updated task and an error if any.
func UpdateTask(newTask Task, id string) (Task, error) {
	objectID, err := ChangeIdToObjectId(id)
	if err != nil {
		return Task{}, err
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

	result, err := TaskCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": updateFields},
	)

	if err != nil {
		return Task{}, errors.New("unable to update Task")
	}

	fmt.Println(result)

	if result.ModifiedCount == 0 {
		return Task{}, errors.New("no task found with the given ID")
	}

	var updatedTask Task
	err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&updatedTask)
	if err != nil {
		return Task{}, err
	}

	return updatedTask, nil
}

// GetTask retrieves a task from the task manager by its ID.
// It takes an ID string as input and returns a Task and an error.
// If the task is found, the Task is populated with the task details.
// If the task is not found, it returns an empty Task and an error indicating that no task was found with the given ID.
func GetTask(id string) (Task, error) {
	objectID, err := ChangeIdToObjectId(id)
	if err != nil {
		return Task{}, err
	}

	var task Task

	result := TaskCollection.FindOne(context.TODO(), bson.M{"_id": objectID})

	if err := result.Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return Task{}, errors.New("no task found with the given ID")
		}
	}

	return task, nil
}

// RemoveTask removes a task from the task manager based on the given ID.
// It takes the ID of the task as a parameter and returns an error if any.
// If the task with the given ID is not found, it returns an error with the message "no task found with the given ID".
// If the task is successfully deleted, it returns nil.
func RemoveTask(id string) error {
	objectID, err := ChangeIdToObjectId(id)
	if err != nil {
		return err
	}

	response, err := TaskCollection.DeleteOne(context.TODO(), bson.M{"_id": objectID})

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
