package middleware

import (
	"net/http"
	"strings"

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
		roleStr, _ := userRole.(string)

		if strings.ToUpper(roleStr) != "ADMIN" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not allowed"})
			c.Abort()
			return
		}

		c.Next()
	}
}
