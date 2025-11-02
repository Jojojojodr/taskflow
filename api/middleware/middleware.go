package middleware

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authentication logic goes here
		// For example, check for a valid JWT token in the request header

		// If authentication fails:
		// c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
		// return

		// If authentication succeeds, proceed to the next handler
		c.Next()
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// User context setup logic goes here
		// For example, extract user information from the JWT token
		// and set it in the context for downstream handlers to use

		// If user context setup fails:
		// c.AbortWithStatusJSON(400, gin.H{"error": "Bad Request"})
		// return

		// If user context setup succeeds, proceed to the next handler
		c.Next()
	}
}