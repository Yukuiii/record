package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/controllers"
	"github.com/sakura/record-api/middleware"
)

// SetupRouter 配置所有路由
func SetupRouter(r *gin.Engine) {
	// 全局中间件
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.RequestLogger())
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimit())

	// 设置404和405处理器
	r.NoRoute(middleware.NotFoundHandler())
	r.NoMethod(middleware.MethodNotAllowedHandler())

	// 健康检查路由（无需认证）
	r.GET("/health", controllers.HealthCheck)
	r.GET("/ping", controllers.Ping)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 无需认证的路由
		auth := v1.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}

		// 需要认证的路由
		user := v1.Group("/user")
		user.Use(middleware.JWTAuth())
		{
			user.GET("/profile", controllers.GetUserProfile)
			user.PUT("/profile", controllers.UpdateUserProfile)
		}

		// 分类相关路由
		categories := v1.Group("/categories")
		categories.Use(middleware.JWTAuth())
		{
			categories.GET("", controllers.GetCategories)
			categories.POST("", controllers.CreateCategory)
			categories.PUT("/:id", controllers.UpdateCategory)
			categories.DELETE("/:id", controllers.DeleteCategory)
		}

		// 交易记录相关路由
		transactions := v1.Group("/transactions")
		transactions.Use(middleware.JWTAuth())
		{
			transactions.GET("", controllers.GetTransactions)
			transactions.POST("", controllers.CreateTransaction)
			transactions.GET("/:id", controllers.GetTransactionByID)
			transactions.PUT("/:id", controllers.UpdateTransaction)
			transactions.DELETE("/:id", controllers.DeleteTransaction)
		}

		// 统计分析相关路由
		statistics := v1.Group("/statistics")
		statistics.Use(middleware.JWTAuth())
		{
			statistics.GET("/monthly", controllers.GetMonthlyStatistics)
			statistics.GET("/yearly", controllers.GetYearlyStatistics)
		}
	}
}
