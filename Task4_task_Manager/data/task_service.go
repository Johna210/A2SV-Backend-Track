package data

import (
	"errors"

	"github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager/models"
)

type Task = models.Task

type TaskManager struct {
	Tasks map[int]Task
}

// GetTasks returns a slice of all tasks in the TaskManager.
func (tm *TaskManager) GetTasks() []Task {
	tasks := []Task{}
	for _, task := range tm.Tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

// GetTask retrieves a task from the TaskManager by its ID.
// It returns the task and a nil error if the task is found.
// If the task is not found, it returns an empty Task struct and an error.
func (tm *TaskManager) GetTask(id int) (Task, error) {
	for taskID, task := range tm.Tasks {
		if taskID == id {
			return task, nil
		}
	}

	return Task{}, errors.New("Task Not Found")

}

// AddTask adds a new task to the task manager.
// It takes a newTask of type Task as a parameter and returns the added task and an error (if any).
// If the task ID already exists in the task manager, it returns an empty task and an error indicating that the task ID is already taken.
func (tm *TaskManager) AddTask(newTask Task) (Task, error) {
	id := newTask.ID

	_, found := tm.Tasks[id]

	if found {
		return Task{}, errors.New("Task ID already taken")
	}

	tm.Tasks[id] = newTask

	return newTask, nil

}

// UpdateTask updates the task with the given ID in the TaskManager.
// It takes the ID of the task to be updated and the newTask object containing the updated task details.
// It returns the updated task and an error if the task is not found.
func (tm *TaskManager) UpdateTask(id int, newTask Task) (Task, error) {
	taskToUpdate, found := tm.Tasks[id]

	if !found {
		return Task{}, errors.New("Task Not Found")
	}

	for TaskID := range tm.Tasks {
		if id == TaskID {
			if newTask.Title != "" {
				taskToUpdate.Title = newTask.Title
			}

			if newTask.Description != "" {
				taskToUpdate.Description = newTask.Description
			}

			if newTask.Status != "" {
				taskToUpdate.Status = newTask.Status
			}

			tm.Tasks[id] = taskToUpdate

			return tm.Tasks[id], nil

		}
	}

	return Task{}, errors.New("Task Not Found")

}

// RemoveTask removes a task from the TaskManager by its ID.
// It returns the removed task and an error if the task is not found.
func (tm *TaskManager) RemoveTask(id int) (Task, error) {
	tasktoDelete, found := tm.Tasks[id]

	if !found {
		return Task{}, errors.New("Task Not found")
	}

	delete(tm.Tasks, id)

	return tasktoDelete, nil
}
