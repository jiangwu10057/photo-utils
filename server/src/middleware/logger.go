package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"server/utils/logger"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		requestTime := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		// proto := c.Request.Proto
		ua := c.Request.UserAgent()

		c.Next()
		
		latency := time.Since(requestTime)
		status := c.Writer.Status()
		loginfo := fmt.Sprintf("%d %s %s %d %s %s \"%s\"", requestTime.Unix(), method, path, status, latency, clientIP, ua)
		logger.Info(loginfo)
	}
}
