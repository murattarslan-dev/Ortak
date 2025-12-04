package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}





// Helper methods for handlers
func SetSuccess(c *gin.Context, message string, data interface{}) {
	c.Set("response_success", true)
	c.Set("response_status", http.StatusOK)
	c.Set("response_message", message)
	c.Set("response_data", data)
}

func SetCreated(c *gin.Context, message string, data interface{}) {
	c.Set("response_success", true)
	c.Set("response_status", http.StatusCreated)
	c.Set("response_message", message)
	c.Set("response_data", data)
}

func SetError(c *gin.Context, statusCode int, message string) {
	c.Set("response_success", false)
	c.Set("response_status", statusCode)
	c.Set("response_message", message)
	c.Abort()
}