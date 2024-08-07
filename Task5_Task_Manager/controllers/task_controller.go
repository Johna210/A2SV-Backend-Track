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
type TaskResponse = models.TaskResponse

var taskManager *data.TaskManager

func InitTaskManagerController(uri string, dbName string, collectionName string) error {
	manager, err := taskManager.ConnectToMongoDB(uri, dbName, collectionName)
	if err != nil {
		return errors.New("couldn't connect to database")
	}
	taskManager = manager

	log.Print("Connected To Database")

	return nil
}

func CloseDB() error {
	if err := taskManager.Client.Disconnect(context.Background()); err != nil {
		return err
	}

	log.Print("Database Closed")

	return nil
}

func GetTasksController(c *gin.Context) {
	tasks, err := taskManager.GetTasks()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Something went wrong"})
	}

	c.JSON(http.StatusOK, tasks)
}

func AddTaskController(c *gin.Context) {
	var newTask Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	if newTask.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Task Title"})
		return
	}

	if newTask.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Task Description"})
		return
	}

	if newTask.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing Task Status"})
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

	c.JSON(http.StatusOK, nil)
}
