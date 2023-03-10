package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	mongoDb "gin-mongo/src/mongoDb"
)

var timeFormat = TimeFormat

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	// claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().AddDate(0, 0, 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func IsValidToken(ctx context.Context, token string) bool {
	err := mongoDb.GetToken(ctx, token)

	if err != nil {
		return false
	}

	return true
}

func ExtractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenId(c *gin.Context) (string, error) {
	tokenString := ExtractToken(c)
	// log.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		mongoDb.DeleteToken(c.Request.Context(), tokenString)
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id := fmt.Sprintf("%v", claims["userId"])
		return id, nil
	}

	return "", nil
}
