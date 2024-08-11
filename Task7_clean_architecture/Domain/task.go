package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStaus string

const (
	NotStarted TaskStaus = "Not Started"
	Started    TaskStaus = "Started"
	Completed  TaskStaus = "Completed"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	Status      TaskStaus          `json:"status" bson:"status" validate:"required,eq=Not Started|eq=Started|eq=Completed"`
	Due_Date    time.Time          `json:"due_date" bson:"due_date"`
}

type TaskRepository interface {
	CreateTask(c context.Context, task *Task) (Task, error)
	Fetch(c context.Context) ([]Task, error)
	GetByID(c context.Context, id primitive.ObjectID) (Task, error)
	UpdateTask(c context.Context, task *Task) (Task, error)
	DeleteTask(c context.Context, id primitive.ObjectID) error
}
