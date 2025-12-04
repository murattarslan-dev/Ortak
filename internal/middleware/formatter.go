package middleware

import (
	"net/http"
	"ortak/pkg/response"

	"github.com/gin-gonic/gin"
)

// FormatterMiddleware processes context data and sends formatted response
func FormatterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// If response already written or error handled, skip
		if c.Writer.Written() {
			return
		}
		if _, exists := c.Get("error_handled"); exists {
			return
		}

		// Check for error response
		if success, exists := c.Get("response_success"); exists && !success.(bool) {
			status := c.GetInt("response_status")
			message := c.GetString("response_message")
			if status == 0 {
				status = http.StatusBadRequest
			}
			c.JSON(status, response.Response{
				Success: false,
				Message: message,
			})
			return
		}

		// Check for success response
		if success, exists := c.Get("response_success"); exists && success.(bool) {
			status := c.GetInt("response_status")
			message := c.GetString("response_message")
			data, _ := c.Get("response_data")
			if status == 0 {
				status = http.StatusOK
			}
			c.JSON(status, response.Response{
				Success: true,
				Message: message,
				Data:    data,
			})
			return
		}

		// If no response was set, return default success
		if _, exists := c.Get("response_success"); !exists {
			c.JSON(http.StatusOK, response.Response{
				Success: true,
				Message: "Success",
			})
		}
	}
}