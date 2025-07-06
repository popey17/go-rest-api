package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/popey17/go-rest-api/utils"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()

}
