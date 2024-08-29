package middleware

import (
	"strings"

	"github.com/Ryan-Albuquerque/go-api/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{
				"message": "Missing Authorization header",
			})
			c.Abort()
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.JSON(401, gin.H{
				"message": "Invalid Authorization header format",
			})
			c.Abort()
			return
		}

		_, err := utils.Verify(token)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
