package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/popey17/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvent)
	server.GET("/events/:id", getEventById)

	authRoutes := server.Group("/")
	authRoutes.Use(middlewares.Auth)
	authRoutes.POST("/events", saveEvent)
	authRoutes.PUT("/events/:id", updateEvent)
	authRoutes.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", saveUser)
	server.POST("/login", login)
}
