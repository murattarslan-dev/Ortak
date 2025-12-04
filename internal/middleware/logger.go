package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

// LoggerMiddleware provides detailed request/response logging
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Read request body
		var requestBody string
		if c.Request.Body != nil {
			bodyBytes, _ := io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			requestBody = string(bodyBytes)
		}

		// Wrap response writer to capture response
		responseBuffer := &bytes.Buffer{}
		w := &responseWriter{
			ResponseWriter: c.Writer,
			body:           responseBuffer,
		}
		c.Writer = w

		c.Next()

		// Log request and response
		latency := time.Since(start)
		responseBody := responseBuffer.String()

		// If no response captured but status is not 200, it might be from middleware
		if responseBody == "" && c.Writer.Status() != 200 {
			responseBody = fmt.Sprintf(`{"status":%d}`, c.Writer.Status())
		}

		// Format request body for logging
		reqLog := "empty"
		if requestBody != "" {
			if len(requestBody) > 200 {
				reqLog = requestBody[:200] + "..."
			} else {
				reqLog = requestBody
			}
			reqLog = strings.ReplaceAll(reqLog, "\n", " ")
		}

		// Format response body for logging
		respLog := "empty"
		if responseBody != "" {
			if len(responseBody) > 200 {
				respLog = responseBody[:200] + "..."
			} else {
				respLog = responseBody
			}
			respLog = strings.ReplaceAll(respLog, "\n", " ")
		}

		log.Printf("[%s] %s %s | Status: %d | Latency: %v | IP: %s\nRequest: %s\nResponse: %s\n",
			start.Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
			c.ClientIP(),
			reqLog,
			respLog,
		)
	}
}
