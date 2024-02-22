package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gabriel-leone/go-jwt/initializers"
	"github.com/gabriel-leone/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Missing authorization cookie",
		})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.JSON(401, gin.H{
			"error": "Error parsing JWT: " + err.Error(),
		})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{
				"error": "JWT is expired",
			})
			c.Abort()
			return
		}

		var user models.User
		result := initializers.DB.First(&user, claims["sub"])

		if result.Error != nil {
			c.JSON(500, gin.H{
				"error": "Database error: " + result.Error.Error(),
			})
			c.Abort()
			return
		}

		if user.ID == 0 {
			c.JSON(401, gin.H{
				"error": "User not found",
			})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	} else {
		c.JSON(401, gin.H{
			"error": "Invalid JWT",
		})
		c.Abort()
	}
}
