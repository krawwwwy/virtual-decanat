package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// LoggingMiddleware логирует информацию о запросах
func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		method := c.Request.Method

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		if query != "" {
			path = path + "?" + query
		}

		logger.Info("HTTP Request",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("client_ip", clientIP),
			zap.String("user_agent", userAgent),
		)
	}
} 