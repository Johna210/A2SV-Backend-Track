package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager/data"
	"github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager/models"
	"github.com/gin-gonic/gin"
)

type Task = models.Task

var taskManager = data.TaskManager{
	Tasks: make(map[int]Task),
}

// GetTasksController is a controller function that handles the GET request to retrieve all tasks.
// It returns a JSON response with the list of tasks.
func GetTasksController(c *gin.Context) {
	c.JSON(http.StatusOK, taskManager.GetTasks())
}

// GetTaskController is a handler function that retrieves a task based on the provided ID.
// It expects the ID to be passed as a URL parameter.
// If the ID is missing or invalid, it returns an appropriate error response.
// If the task is found, it returns the task details in the response body.
func GetTaskController(c *gin.Context) {
	idChar := c.Param("id")
	if idChar == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing param id"})
		return
	}

	id, err := strconv.Atoi(idChar)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid param type"})
		return
	}

	task, err := taskManager.GetTask(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// AddTaskController handles the HTTP POST request to add a new task.
// It binds the JSON request body to a newTask struct, validates it, and adds the task using the taskManager.
// If the request body is invalid or there is an error adding the task, it returns an appropriate JSON response.
// If the task is added successfully, it returns the created task in the response body.
func AddTaskController(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	createdTask, err := taskManager.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTaskController updates a task with the given ID.
// It expects a JSON request body containing the new task details.
// If the ID is invalid or the request body is invalid, it returns a JSON response with an error message.
// If the task is not found, it returns a JSON response with a "Task Not Found" message.
// Otherwise, it returns a JSON response with the updated task.
func UpdateTaskController(c *gin.Context) {
	id, err := idHelper(c)
	var newTask Task

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	updatedTask, err := taskManager.UpdateTask(id, newTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTaskController handles the deletion of a task based on the provided ID.
// It expects the ID to be passed as a query parameter in the request URL.
// If the ID is valid and the task is found, it will be removed from the task manager.
// If the ID is invalid or the task is not found, an appropriate error response will be returned.
func DeleteTaskController(c *gin.Context) {
	id, err := idHelper(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid param id"})
		return
	}

	task, err := taskManager.RemoveTask(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Task Not Found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// Helper function to take id from Param
// idHelper is a helper function that extracts and converts the "id" parameter from the request context.
// It returns the extracted ID as an integer and an error if there is any issue with the parameter.
func idHelper(c *gin.Context) (int, error) {
	idChar := c.Param("id")
	if idChar == "" {
		return 0, fmt.Errorf("missing param id")
	}

	id, err := strconv.Atoi(idChar)
	if err != nil {
		return 0, fmt.Errorf("invalid param type")
	}

	return id, nil
}
