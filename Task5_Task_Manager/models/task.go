package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type TaskStatus string

const (
	NotStarted TaskStatus = "Not Started"
	Started    TaskStatus = "Started"
	Completed  TaskStatus = "Completed"
)

// Task struct for creating a new Task
type Task struct {
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Status      TaskStatus `json:"status" bson:"status"`
}

// TaskResponse struct for responses, including the ID
type TaskResponse struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      TaskStatus         `json:"status" bson:"status"`
}
