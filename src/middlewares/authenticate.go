package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// log.Println(authHeader)

		if authHeader == "12456" {
			c.Next()
		} else {
			// log.Println("Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "unauthorized"})
		}
	}
}
