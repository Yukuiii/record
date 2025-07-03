package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/services"
)

// 创建统计分析服务实例
var statisticsService = services.NewStatisticsService()

// GetMonthlyStatistics 获取月度统计
func GetMonthlyStatistics(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 绑定查询参数
	var params api.MonthlyStatisticsQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的查询参数: "+err.Error()))
		return
	}

	// 调用服务层获取月度统计
	statistics, err := statisticsService.GetMonthlyStatistics(userID.(uint), params.Year, params.Month)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"statistics": statistics,
	}))
}

// GetYearlyStatistics 获取年度统计
func GetYearlyStatistics(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 绑定查询参数
	var params api.YearlyStatisticsQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的查询参数: "+err.Error()))
		return
	}

	// 调用服务层获取年度统计
	statistics, err := statisticsService.GetYearlyStatistics(userID.(uint), params.Year)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"statistics": statistics,
	}))
}