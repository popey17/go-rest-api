package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/popey17/go-rest-api/db"
	"github.com/popey17/go-rest-api/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvent)
	server.POST("/events", saveEvent)

	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	events := models.GetAll()
	context.JSON(http.StatusOK, events)
}

func saveEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	event.ID = 1
	event.UserID = 1

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}
