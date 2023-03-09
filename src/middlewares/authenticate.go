package middleware

import (
	"net/http"

	token "gin-mongo/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := token.ExtractToken(c)
		r := c.Request
		// log.Println(authHeader)
		if token.IsValidToken(r.Context(), accessToken) {
			c.Next()
		} else {
			// log.Println("Unauthorized")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": 401, "message": "unauthorized"})
		}
	}
}
