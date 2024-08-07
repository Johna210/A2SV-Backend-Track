package main

import (
	routes "github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager/router"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Routes(router)

	router.Run(":4000")

}
