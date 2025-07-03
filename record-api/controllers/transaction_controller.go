package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/services"
)

// 创建交易记录服务实例
var transactionService = services.NewTransactionService()

// GetTransactions 获取交易记录列表
func GetTransactions(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 绑定查询参数
	var params api.TransactionQueryParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的查询参数: "+err.Error()))
		return
	}

	// 调用服务层获取交易记录
	transactions, total, err := transactionService.GetTransactions(userID.(uint), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// 计算分页信息
	totalPages := (int(total) + params.PageSize - 1) / params.PageSize

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"transactions": transactions,
		"pagination": gin.H{
			"page":        params.Page,
			"page_size":   params.PageSize,
			"total":       total,
			"total_pages": totalPages,
		},
	}))
}

// CreateTransaction 创建交易记录
func CreateTransaction(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 绑定请求参数
	var req api.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层创建交易记录
	transaction, err := transactionService.CreateTransaction(userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusCreated, api.SuccessResponse(gin.H{
		"transaction": transaction,
	}))
}

// GetTransactionByID 根据ID获取交易记录
func GetTransactionByID(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 获取交易记录ID
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的交易记录ID"))
		return
	}

	// 调用服务层获取交易记录
	transaction, err := transactionService.GetTransactionByID(userID.(uint), uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"transaction": transaction,
	}))
}

// UpdateTransaction 更新交易记录
func UpdateTransaction(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 获取交易记录ID
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的交易记录ID"))
		return
	}

	// 绑定请求参数
	var req api.TransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层更新交易记录
	transaction, err := transactionService.UpdateTransaction(userID.(uint), uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"transaction": transaction,
	}))
}

// DeleteTransaction 删除交易记录
func DeleteTransaction(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 获取交易记录ID
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的交易记录ID"))
		return
	}

	// 调用服务层删除交易记录
	if err := transactionService.DeleteTransaction(userID.(uint), uint(id)); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"message": "交易记录删除成功",
	}))
}