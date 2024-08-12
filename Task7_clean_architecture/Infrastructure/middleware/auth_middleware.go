package middleware

import (
	"log"
	"net/http"
	"strings"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Bootstrap"
	infrastructure "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Infrastructure"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware function that handles authentication for incoming requests.
// It checks for the presence of an Authorization header and validates the JWT token.
// If the token is valid, it extracts the user role from the token claims and sets it in the context.
// If the token is invalid or the user role is not allowed, it returns an error response and aborts the request.
// This middleware should be used to protect routes that require authentication.
func AuthMiddleware(env *bootstrap.Env) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// check if there is Authorization Header
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// check if there is bearer sent also
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		err := infrastructure.IsAuthorized(authParts[1], string(env.AccessTokenSecret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		authorized, err := infrastructure.CheckTokenExpiry(authParts[1], string(env.AccessTokenSecret))
		if !authorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		claims, err := infrastructure.ExtractClaims(authParts[1], string(env.AccessTokenSecret))
		log.Printf("claims is here %v", claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		userID := claims["UserID"]
		userRole := claims["Role"]

		c.Set("userID", userID)
		c.Set("userRole", userRole)
		c.Set("claims", claims)

		c.Next()
	}
}
