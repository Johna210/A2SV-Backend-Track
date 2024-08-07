package main

import (
	"log"

	"github.com/Johna210/A2SV-Backend-Track/Track5_Task_Manager/controllers"
	routes "github.com/Johna210/A2SV-Backend-Track/Track5_Task_Manager/router"
	"github.com/gin-gonic/gin"
)

func main() {
	uri := "mongodb://localhost:27017"
	dbName := "taskManager"
	collectionName := "tasks"

	// Connect with DB
	err := controllers.InitTaskManagerController(uri, dbName, collectionName)

	if err != nil {
		log.Fatal(err)
	}

	// Close DB after finished
	defer func() {
		if err = controllers.CloseDB(); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	routes.Routes(router)

	router.Run(":4000")
}
