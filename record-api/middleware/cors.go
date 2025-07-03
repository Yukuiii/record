package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/config"
)

// CORS 跨域资源共享中间件
func CORS() gin.HandlerFunc {
	// 获取配置
	cfg := config.GetConfig()

	// 配置CORS
	corsConfig := cors.DefaultConfig()

	// 根据运行模式设置不同的CORS策略
	if cfg.Server.Mode == "debug" {
		// 开发模式：允许所有来源
		corsConfig.AllowAllOrigins = true
		corsConfig.AllowCredentials = true
	} else {
		// 生产模式：只允许特定来源
		corsConfig.AllowOrigins = []string{
			"https://yourdomain.com",
			"https://www.yourdomain.com",
		}
		corsConfig.AllowCredentials = true
	}

	// 允许的HTTP方法
	corsConfig.AllowMethods = []string{
		"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH",
	}

	// 允许的请求头
	corsConfig.AllowHeaders = []string{
		"Origin", "Content-Length", "Content-Type", "Authorization",
		"X-Requested-With", "Accept", "Accept-Encoding", "Accept-Language",
	}

	// 暴露的响应头
	corsConfig.ExposeHeaders = []string{
		"Content-Length", "Content-Type",
	}

	return cors.New(corsConfig)
}