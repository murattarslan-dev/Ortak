package middleware

import (
	"net/http"
	"ortak/pkg/response"
	"ortak/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Authorization header required",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		storage := utils.GetMemoryStorage()
		userID, exists := storage.IsTokenValid(tokenString)
		if !exists || userID != claims.UserID {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Token not found or invalid",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}