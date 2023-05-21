package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: JWT Authorization; current implementaion is temporary
// Authenticate extracts a user from the Authorization header
// It sets the user to the context if the user exists
func Authenticate(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		if authHeader != "Bearer "+token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			return
		}

		c.Next()
	}
}
