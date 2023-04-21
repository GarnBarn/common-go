package httpserver

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthModelMapping() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader(UserUidKey)
		splittedKey := strings.Split(key, " ")
		if len(splittedKey) != 2 {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		tokenStr := splittedKey[1]

		token, err := jwt.Parse(tokenStr, nil)
		if token == nil || err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		result, ok := claims["user_id"].(string)

		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Header(UserUidKey, result)
		c.Next()
	}
}
