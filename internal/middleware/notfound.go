package middleware

import (
	"fmt"
	"net/http"
	"ortak/pkg/response"

	"github.com/gin-gonic/gin"
)

// NotFoundMiddleware handles 404 errors with standard response format
func NotFoundMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := fmt.Sprintf("Endpoint not found: %s %s", c.Request.Method, c.Request.URL.Path)
		c.JSON(http.StatusNotFound, response.Response{
			Success: false,
			Message: message,
		})
	}
}