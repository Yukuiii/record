package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/services"
)

// 创建认证服务实例
var authService = services.NewAuthService()

// Register 用户注册控制器
func Register(c *gin.Context) {
	// 绑定请求参数
	var req api.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层进行注册
	userID, token, err := authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"user_id": userID,
		"token":   token,
	}))
}

// Login 用户登录控制器
func Login(c *gin.Context) {
	// 绑定请求参数
	var req api.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层进行登录
	user, token, err := authService.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"user": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"phone":    user.Phone,
			"nickname": user.Nickname,
			"avatar":   user.Avatar,
		},
		"token": token,
	}))
}
