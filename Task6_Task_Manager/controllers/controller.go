package controllers

import (
	"fmt"
	"net/http"

	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/data"
	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Task = models.Task

var validate = validator.New()

// RegisterUser is a function that handles the registration of a user.
// It takes a gin.Context object as a parameter and expects the user information to be provided in the request body as JSON.
// The function first binds the JSON data to a models.User struct and performs validation on the user data.
// If the data is valid, it calls the data.Register function to register the user.
// The function returns a JSON response with the registered user's data and a success message if the registration is successful.
// If there are any errors during the registration process, appropriate error messages are returned in the JSON response.
func RegisterUser(c *gin.Context) {
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

	returnValue, err := data.Register(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully", "data": returnValue})
}

// LoginUser handles the login functionality for a user.
// It expects a JSON payload containing the user's username and password.
// If the payload is valid, it attempts to authenticate the user and generate a token.
// If successful, it returns the generated token in the response body.
// If there are any errors during the process, it returns an error message in the response body.
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

// Promote promotes a user identified by the given ID.
// It retrieves the ID from the request context and calls the data.Promote function to perform the promotion.
// If an error occurs while retrieving the ID or promoting the user, it returns a JSON response with the corresponding error message.
// If the promotion is successful, it returns a JSON response with the message "user promoted" and a status code of 200.
func Promote(c *gin.Context) {
	id, err := idHelper(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = data.Promote(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user promoted"})
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
