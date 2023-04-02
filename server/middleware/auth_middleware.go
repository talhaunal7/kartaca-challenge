package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"example.com/auction-api/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(redis service.RedisService) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("Authorization")
		fmt.Println("TOKEN", tokenString)
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Couldn't get token"})
			return

		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			log.Print(err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is expired"})
				return
			}
		}
		userId := claims["sub"]
		redisToken, err := redis.Get(fmt.Sprintf("%v", userId))

		if redisToken == nil || *redisToken != tokenString || err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("id", userId)
		c.Next()

	}
}
