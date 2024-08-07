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

func GetTasksController(c *gin.Context) {
	c.JSON(http.StatusOK, taskManager.GetTasks())
}

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
