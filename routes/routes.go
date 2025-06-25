package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvent)
	server.POST("/events", saveEvent)

	server.GET("/events/:id", getEventById)

	server.PUT("/events/:id", updateEvent)

	server.DELETE("/events/:id", deleteEvent)
}
