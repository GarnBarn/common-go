package httpserver

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthModelMapping() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUserId := c.GetHeader(UserUidKey)
		if currentUserId != "" {
			c.Next()
			return
		}

		key := c.GetHeader("Authorization")

		if key == "" {
			key = c.GetHeader("authorization")
		}

		splittedKey := strings.Split(key, " ")
		if len(splittedKey) != 2 {
			logrus.Warn("Key not correctly structed.")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		tokenStr := splittedKey[1]

		logrus.Info("Token Str: ", tokenStr)

		token, _ := jwt.Parse(tokenStr, nil)
		if token == nil {
			logrus.Warn("Can't parse jwt")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		claims, _ := token.Claims.(jwt.MapClaims)

		result, ok := claims["user_id"].(string)

		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Request.Header.Add(UserUidKey, result)
		c.Next()
	}
}
