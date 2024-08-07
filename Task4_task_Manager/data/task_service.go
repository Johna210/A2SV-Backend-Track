package data

import (
	"errors"

	"github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager/models"
)

type Task = models.Task

type TaskManager struct {
	Tasks map[int]Task
}

func (tm *TaskManager) GetTasks() []Task {
	tasks := []Task{}
	for _, task := range tm.Tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

func (tm *TaskManager) GetTask(id int) (Task, error) {
	for taskID, task := range tm.Tasks {
		if taskID == id {
			return task, nil
		}
	}

	return Task{}, errors.New("Task Not Found")

}

func (tm *TaskManager) AddTask(newTask Task) (Task, error) {
	id := newTask.ID

	_, found := tm.Tasks[id]

	if found {
		return Task{}, errors.New("Task ID already taken")
	}

	tm.Tasks[id] = newTask

	return newTask, nil

}

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

func (tm *TaskManager) RemoveTask(id int) (Task, error) {
	tasktoDelete, found := tm.Tasks[id]

	if !found {
		return Task{}, errors.New("Task Not found")
	}

	delete(tm.Tasks, id)

	return tasktoDelete, nil
}
