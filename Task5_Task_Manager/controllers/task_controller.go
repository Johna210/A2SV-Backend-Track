package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/Johna210/A2SV-Backend-Track/Track5_Task_Manager/data"
	"github.com/Johna210/A2SV-Backend-Track/Track5_Task_Manager/models"
	"github.com/gin-gonic/gin"
)

type Task = models.Task

var taskManager *data.TaskManager

// InitTaskManagerController initializes the task manager controller by connecting to the MongoDB database.
// It takes the MongoDB connection URI, database name, and collection name as parameters.
// Returns an error if the connection to the database fails.
func InitTaskManagerController(uri string, dbName string, collectionName string) error {
	manager, err := taskManager.ConnectToMongoDB(uri, dbName, collectionName)
	if err != nil {
		return errors.New("couldn't connect to database")
	}
	taskManager = manager

	log.Print("Connected To Database")

	return nil
}

// CloseDB closes the database connection.
// It disconnects the client from the database and logs a message when the database is closed.
// Returns an error if there was a problem disconnecting from the database.
func CloseDB() error {
	if err := taskManager.Client.Disconnect(context.Background()); err != nil {
		return err
	}

	log.Print("Database Closed")

	return nil
}

// GetTasksController handles the HTTP GET request to retrieve all tasks.
// It calls the GetTasks function from the taskManager package to fetch the tasks.
// If an error occurs, it returns a JSON response with a status code of 404.
// Otherwise, it returns a JSON response with a status code of 200 and the tasks data.
func GetTasksController(c *gin.Context) {
	tasks, err := taskManager.GetTasks()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
	}

	c.JSON(http.StatusOK, tasks)
}

// AddTaskController handles the HTTP POST request to add a new task.
// It binds the JSON request body to the `newTask` variable and performs validation checks.
// If the request body is invalid or any required fields are missing, it returns an appropriate error response.
// Otherwise, it calls the `AddTask` function of the `taskManager` to add the task.
// If the task is added successfully, it returns the created task in the response.
func AddTaskController(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body."})
		return
	}

	if newTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Task Title."})
		return
	}

	if newTask.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Task Description."})
		return
	}

	if newTask.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Task Status."})
		return
	}

	if newTask.DueDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Due Date."})
		return
	}

	createdTask, err := taskManager.AddTask(newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

// Helper function to take id from Param
// idHelper is a helper function that extracts and converts the "id" parameter from the request context.
// It returns the extracted ID as an integer and an error if there is any issue with the parameter.
func idHelper(c *gin.Context) (string, error) {
	idChar := c.Param("id")
	if idChar == "" {
		return "", fmt.Errorf("missing param id")
	}

	return idChar, nil
}

// UpdateTaskController is a controller function that handles the update of a task.
// It receives a request context `c` from the Gin framework and updates the task with the specified ID.
// The updated task is returned as a JSON response.
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

	updatedTask, err := taskManager.UpdateTask(newTask, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// GetTaskController handles the HTTP GET request to retrieve a task by its ID.
// It expects the task ID to be provided as a query parameter.
// If the ID is valid and the task is found, it returns the task details in the response body.
// If the ID is invalid or the task is not found, it returns an appropriate error message.
func GetTaskController(c *gin.Context) {
	id, err := idHelper(c)

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

// DeleteTaskController is a controller function that handles the deletion of a task.
// It expects a valid task ID as a parameter in the request and removes the task from the task manager.
// If the task ID is invalid or the task is not found, it returns an appropriate error response.
// If the task is successfully removed, it returns a success message.
func DeleteTaskController(c *gin.Context) {
	id, err := idHelper(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid param id"})
		return
	}

	err = taskManager.RemoveTask(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Task Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task Removed successfully."})
}
