package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		done := make(chan struct{})
		go func() {
			select {
			case <-done:
			case <-time.After(d):
				c.AbortWithStatus(http.StatusRequestTimeout)
				return
			}
		}()
		c.Next()
		close(done)
	}
}
