package main

import (
	"github.com/gin-gonic/gin"
	"github.com/popey17/go-rest-api/db"
	"github.com/popey17/go-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
