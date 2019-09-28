package middleware

import (
	"github.com/gin-gonic/gin"
	"time"

	log "github.com/cihub/seelog"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Info(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Info(status)
	}
}
