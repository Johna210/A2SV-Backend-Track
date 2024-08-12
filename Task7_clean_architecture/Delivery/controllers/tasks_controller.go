package controllers

import (
	"net/http"

	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase domain.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid task format"})
		return
	}

	task.ID = primitive.NewObjectID()
	err = tc.TaskUsecase.CreateTask(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (tc *TaskController) Fetch(c *gin.Context) {
	tasks, err := tc.TaskUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task domain.TaskUpdate
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid task format"})
		return
	}

	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id param missing"})
		return
	}

	updatedTask, err := tc.TaskUsecase.UpdateTask(c, &task, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, updatedTask)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id param missing"})
		return
	}

	task, err := tc.TaskUsecase.GetByID(c, taskID)

	if task.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "task with the given id not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id param missing"})
		return
	}

	err := tc.TaskUsecase.DeleteTask(c, taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task removed successfully."})
}
