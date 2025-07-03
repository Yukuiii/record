package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/database"
)

// HealthCheck 健康检查接口
func HealthCheck(c *gin.Context) {
	// 检查数据库连接
	dbStatus := "ok"
	var dbInfo map[string]interface{}

	if err := database.HealthCheck(); err != nil {
		dbStatus = "error: " + err.Error()
	} else {
		dbInfo = database.GetDatabaseInfo()
	}

	// 返回健康状态
	response := gin.H{
		"status":    "ok",
		"timestamp": time.Now().Unix(),
		"services": gin.H{
			"database": dbStatus,
		},
		"version": "1.0.0",
	}

	// 如果数据库连接正常，添加详细信息
	if dbStatus == "ok" && dbInfo != nil {
		response["database_info"] = dbInfo
	}

	c.JSON(http.StatusOK, api.SuccessResponse(response))
}

// Ping 简单的ping接口
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"message": "pong",
		"timestamp": time.Now().Unix(),
	}))
}
