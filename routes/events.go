package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/popey17/go-rest-api/models"
)

func getEvent(context *gin.Context) {
	events, err := models.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again"})
		return
	}
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

	err = event.Save()
	if err != err {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save events. Try again"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func getEventById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event for this Id is not found."})
		return
	}
	result, err := models.GetById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Fetch successful", "event": result})

}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event for this Id is not found."})
		return
	}

	_, err = models.GetById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	updatedEvent.ID = id

	err = updatedEvent.UpdateEvent()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event. Try again"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Updated", "event": updatedEvent})

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Event for this Id is not found."})
		return
	}
	result, err := models.GetById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}

	err = result.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event. Try again"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Deleted", "event": result})

}
