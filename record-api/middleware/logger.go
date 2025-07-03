package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/utils"
)

// RequestLogger 请求日志中间件
func RequestLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 记录请求信息
		utils.LogInfof("[%s] %s %s %d %s %s",
			param.TimeStamp.Format("2006-01-02 15:04:05"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
		)
		return ""
	})
}

// DetailedRequestLogger 详细请求日志中间件（包含请求体）
func DetailedRequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录开始时间
		start := time.Now()

		// 读取请求体
		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			// 重新设置请求体，以便后续处理
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// 记录请求信息
		utils.LogDebugf("Request: %s %s from %s, Body: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			string(requestBody),
		)

		// 处理请求
		c.Next()

		// 记录响应信息
		latency := time.Since(start)
		utils.LogDebugf("Response: %d in %v for %s %s",
			c.Writer.Status(),
			latency,
			c.Request.Method,
			c.Request.URL.Path,
		)
	}
}
