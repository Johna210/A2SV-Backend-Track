package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not allowed"})
			c.Abort()
			return
		}

		userRole, ok := c.Get("userRole")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not allowed"})
			c.Abort()
			return
		}

		if userRole == "USER" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not allowed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
