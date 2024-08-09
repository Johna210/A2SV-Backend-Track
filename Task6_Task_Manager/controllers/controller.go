package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/data"
	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
)

type Task = models.Task

var validate = validator.New()

var userCollection = data.DB.UserCollection

func RegisterUser(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}
	// Check if email already taken
	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for email"})
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
		return
	}

	// Check if username already taken
	count, err = userCollection.CountDocuments(ctx, bson.M{"user_name": user.User_Name})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for email"})
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user name already taken"})
		return
	}

	// Check if there is no any user in the data base and make the user admin else normal user.
	count, err = userCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for email"})
	}

	if count == 0 {
		adminRole := "ADMIN"
		user.User_Role = &adminRole
	}

	returnValue, err := data.Register(*user.First_Name, *user.Last_Name, *user.User_Name,
		*user.Email, *user.Password, *user.User_Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("unable to register user: %v", err.Error())})
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully", "data": returnValue})
}

func LoginUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := data.Login(*user.User_Name, *user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetTasksController handles the HTTP GET request to retrieve all tasks.
// It calls the GetTasks function from the taskManager package to fetch the tasks.
// If an error occurs, it returns a JSON response with a status code of 404.
// Otherwise, it returns a JSON response with a status code of 200 and the tasks data.
func GetTasksController(c *gin.Context) {
	tasks, err := data.GetTasks()

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

	createdTask, err := data.AddTask(newTask)
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

	updatedTask, err := data.UpdateTask(newTask, id)
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

	task, err := data.GetTask(id)

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

	err = data.RemoveTask(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Task Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task Removed successfully."})
}
