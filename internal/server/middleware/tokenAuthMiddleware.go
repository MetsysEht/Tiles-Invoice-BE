package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/MetsysEht/Tiles-Invoice-BE/internal/boot"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/database/models"
	"github.com/MetsysEht/Tiles-Invoice-BE/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var OpenRoutes = []string{
	"/users/login",
}

func CheckAuthMiddleware(c *gin.Context) {
	for _, route := range OpenRoutes {
		if strings.Contains(c.Request.URL.Path, route) {
			c.Next()
			return
		}
	}
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(boot.Config.App.Key), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var user models.User
	boot.DB.Where("username=?", claims["sub"]).Find(&user)

	if utils.IsEmpty(user.ID) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("username", user.Username)
	c.Set("role", user.Role)

	c.Next()
}
