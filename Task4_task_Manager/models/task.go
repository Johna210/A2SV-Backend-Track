package models

type TaskStatus string

const (
	NotStarted TaskStatus = "Not Started"
	Started    TaskStatus = "Started"
	Completed  TaskStatus = "Completed"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}
