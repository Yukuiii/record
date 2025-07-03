package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
)

// 简单的内存限流实现
type limiter struct {
	mu      sync.Mutex
	records map[string][]time.Time
	window  time.Duration
	max     int
}

// 创建新的限流器
func newLimiter(window time.Duration, max int) *limiter {
	return &limiter{
		records: make(map[string][]time.Time),
		window:  window,
		max:     max,
	}
}

// 检查是否允许请求
func (l *limiter) allow(key string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-l.window)

	// 清理过期记录
	var validRecords []time.Time
	for _, t := range l.records[key] {
		if t.After(windowStart) {
			validRecords = append(validRecords, t)
		}
	}

	l.records[key] = validRecords

	// 检查是否超过限制
	if len(validRecords) >= l.max {
		return false
	}

	// 添加新的请求记录
	l.records[key] = append(l.records[key], now)
	return true
}

// 全局限流器实例
var rateLimiter = newLimiter(time.Minute, 60) // 每分钟60个请求

// RateLimit 请求频率限制中间件
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 使用客户端IP作为限流键
		clientIP := c.ClientIP()

		// 检查是否允许请求
		if !rateLimiter.allow(clientIP) {
			c.JSON(http.StatusTooManyRequests, api.Response{
				Code:    http.StatusTooManyRequests,
				Message: "请求过于频繁，请稍后再试",
				Data:    nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
