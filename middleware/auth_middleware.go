package middleware

import (
	"example.com/auction-api/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

func ValidateToken(redis service.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		tokenString = strings.Split(tokenString, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}

		userId := claims["sub"]
		redisToken, err := redis.Get(fmt.Sprintf("%v", userId))
		if err != nil || *redisToken != tokenString {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("id", userId)
		c.Next()

	}
}

