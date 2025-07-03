package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/utils"
)

// ErrorHandler 错误处理中间件
func ErrorHandler() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		// 记录错误日志
		utils.LogErrorf("Panic recovered: %v\nStack trace:\n%s", recovered, debug.Stack())

		// 返回统一的错误响应
		c.JSON(http.StatusInternalServerError, api.ErrorResponse(
			http.StatusInternalServerError,
			"服务器内部错误，请稍后重试",
		))
	})
}

// NotFoundHandler 404处理中间件
func NotFoundHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, api.ErrorResponse(
			http.StatusNotFound,
			"请求的资源不存在",
		))
	}
}

// MethodNotAllowedHandler 405处理中间件
func MethodNotAllowedHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, api.ErrorResponse(
			http.StatusMethodNotAllowed,
			"请求方法不被允许",
		))
	}
}
