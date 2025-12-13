package middleware

import (
	"database/sql"
	"net/http"
	"ortak/pkg/response"
	"ortak/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(db *sql.DB) gin.HandlerFunc {
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

		// Check token in database
		var userID string
		err = db.QueryRow("SELECT user_id FROM tokens WHERE token = $1", tokenString).Scan(&userID)
		if err != nil || userID != claims.UserID {
			c.JSON(http.StatusUnauthorized, response.Response{
				Success: false,
				Message: "Token not found or invalid",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}