package middleware

import (
	"net/http"
	"ortak/pkg/response"

	"github.com/gin-gonic/gin"
)

// ErrorMiddleware catches and formats unexpected errors
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only handle unexpected errors (not bind errors - those are handled in handlers)
		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			if !c.Writer.Written() {
				switch err.Type {
				case gin.ErrorTypePublic:
					c.JSON(http.StatusBadRequest, response.Response{
						Success: false,
						Message: err.Error(),
					})
				default:
					c.JSON(http.StatusInternalServerError, response.Response{
						Success: false,
						Message: "Something went wrong",
					})
				}
			}
		}
	}
}
