package middleware

import (
	"example/restapi/internal/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token doesn't have Bearer prefix"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authorizationHeader, "Bearer ")
		cfg := config.Load()

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(cfg.JWTSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			c.Abort()
			return
		}

		fmt.Println("token claims!!!!!!: ", token.Claims)

		c.Set("Author", token.Claims)

		c.Next()
	}
}
