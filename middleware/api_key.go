package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const apiKey = "my-secret-key"

func APIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")

		if key != apiKey {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid api key",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
