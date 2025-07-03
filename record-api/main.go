package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/config"
	"github.com/sakura/record-api/database"
	"github.com/sakura/record-api/routes"
	"github.com/sakura/record-api/utils"
)

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化日志系统
	if err := utils.InitLogger(); err != nil {
		log.Fatalf("初始化日志系统失败: %v", err)
	}

	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 设置运行模式
	if config.GetConfig().Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化路由
	r := gin.Default()
	routes.SetupRouter(r)

	// 启动服务器
	port := config.GetConfig().Server.Port
	fmt.Printf("服务器启动在 http://localhost:%s\n", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
