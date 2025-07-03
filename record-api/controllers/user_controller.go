package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/services"
)

// 创建用户服务实例
var userService = services.NewUserService()

// GetUserProfile 获取用户资料
func GetUserProfile(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 调用服务层获取用户
	user, err := userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"user": gin.H{
			"id":           user.ID,
			"email":        user.Email,
			"phone":        user.Phone,
			"nickname":     user.Nickname,
			"avatar":       user.Avatar,
			"gender":       user.Gender,
			"birthday":     user.Birthday,
			"register_time": user.RegisterTime,
			"last_login":   user.LastLogin,
		},
	}))
}

// UpdateUserProfile 更新用户资料
func UpdateUserProfile(c *gin.Context) {
	// 从上下文获取用户ID
	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, "未认证"))
		return
	}

	// 绑定请求参数
	var req api.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "无效的请求参数: "+err.Error()))
		return
	}

	// 调用服务层更新用户
	updatedUser, err := userService.UpdateProfile(userID.(uint), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, api.SuccessResponse(gin.H{
		"user": gin.H{
			"id":       updatedUser.ID,
			"nickname": updatedUser.Nickname,
			"avatar":   updatedUser.Avatar,
			"gender":   updatedUser.Gender,
			"birthday": updatedUser.Birthday,
		},
	}))
}
