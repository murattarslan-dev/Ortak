package middleware

import (
	"net/http"
	"ortak/pkg/response"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware handles panic recovery
func RecoveryMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError, response.Response{
				Success: false,
				Message: "Internal server error: " + err,
			})
		} else {
			c.JSON(http.StatusInternalServerError, response.Response{
				Success: false,
				Message: "Internal server error",
			})
		}
		c.Abort()
	})
}