package middlewares

import (
	"ecommerce-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware to check if a user is authenticated
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, role, err := utils.ExtractTokenData(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Set("role", role)
		c.Next()
	}
}

// Middleware to restrict access to admin users
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admins only"})
			c.Abort()
			return
		}
		c.Next()
	}
}
