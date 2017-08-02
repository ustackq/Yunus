package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ustack/Yunus/src/app/backend/pkg/flowcontroller"
)

// RateLimiter return RateLimiter
func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		qps := float32(c.GetFloat64("QPS"))
		burst := c.GetInt("BURST")
		ratelimiter := flowcontroller.NewTokenBucketRateLimiter(qps, burst)
		if !ratelimiter.TryAccept() {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
	}
}
