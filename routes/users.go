package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/popey17/go-rest-api/models"
)

func saveUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": user})
}
