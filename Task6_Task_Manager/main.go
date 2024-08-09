package main

import (
	"log"

	"github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/data"
	routes "github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/router"
	"github.com/gin-gonic/gin"
)

// func init() {
// }

func main() {
	data.Init()

	// Close DB after finished
	defer func() {
		if err := data.CloseDB(); err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()

	routes.Routes(router)

	router.Run(":4000")
}
