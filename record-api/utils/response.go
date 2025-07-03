package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakura/record-api/api"
)

// ResponseHelper 响应助手结构
type ResponseHelper struct {
	ctx *gin.Context
}

// NewResponseHelper 创建响应助手
func NewResponseHelper(c *gin.Context) *ResponseHelper {
	return &ResponseHelper{ctx: c}
}

// Success 成功响应
func (r *ResponseHelper) Success(data interface{}) {
	r.ctx.JSON(http.StatusOK, api.SuccessResponse(data))
}

// Created 创建成功响应
func (r *ResponseHelper) Created(data interface{}) {
	r.ctx.JSON(http.StatusCreated, api.SuccessResponse(data))
}

// BadRequest 请求错误响应
func (r *ResponseHelper) BadRequest(message string) {
	r.ctx.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, message))
}

// Unauthorized 未授权响应
func (r *ResponseHelper) Unauthorized(message string) {
	if message == "" {
		message = "未授权访问"
	}
	r.ctx.JSON(http.StatusUnauthorized, api.ErrorResponse(http.StatusUnauthorized, message))
}

// Forbidden 禁止访问响应
func (r *ResponseHelper) Forbidden(message string) {
	if message == "" {
		message = "禁止访问"
	}
	r.ctx.JSON(http.StatusForbidden, api.ErrorResponse(http.StatusForbidden, message))
}

// NotFound 未找到响应
func (r *ResponseHelper) NotFound(message string) {
	if message == "" {
		message = "资源未找到"
	}
	r.ctx.JSON(http.StatusNotFound, api.ErrorResponse(http.StatusNotFound, message))
}

// InternalServerError 服务器内部错误响应
func (r *ResponseHelper) InternalServerError(message string) {
	if message == "" {
		message = "服务器内部错误"
	}
	r.ctx.JSON(http.StatusInternalServerError, api.ErrorResponse(http.StatusInternalServerError, message))
}

// TooManyRequests 请求过多响应
func (r *ResponseHelper) TooManyRequests(message string) {
	if message == "" {
		message = "请求过于频繁，请稍后再试"
	}
	r.ctx.JSON(http.StatusTooManyRequests, api.ErrorResponse(http.StatusTooManyRequests, message))
}

// ValidationError 验证错误响应
func (r *ResponseHelper) ValidationError(message string) {
	r.ctx.JSON(http.StatusBadRequest, api.ErrorResponse(http.StatusBadRequest, "验证失败: "+message))
}

// CustomError 自定义错误响应
func (r *ResponseHelper) CustomError(code int, message string) {
	r.ctx.JSON(code, api.ErrorResponse(code, message))
}

// 全局响应函数

// SuccessResponse 全局成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	NewResponseHelper(c).Success(data)
}

// CreatedResponse 全局创建成功响应
func CreatedResponse(c *gin.Context, data interface{}) {
	NewResponseHelper(c).Created(data)
}

// ErrorResponse 全局错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	NewResponseHelper(c).CustomError(code, message)
}

// BadRequestResponse 全局请求错误响应
func BadRequestResponse(c *gin.Context, message string) {
	NewResponseHelper(c).BadRequest(message)
}

// UnauthorizedResponse 全局未授权响应
func UnauthorizedResponse(c *gin.Context, message string) {
	NewResponseHelper(c).Unauthorized(message)
}

// NotFoundResponse 全局未找到响应
func NotFoundResponse(c *gin.Context, message string) {
	NewResponseHelper(c).NotFound(message)
}

// InternalServerErrorResponse 全局服务器错误响应
func InternalServerErrorResponse(c *gin.Context, message string) {
	NewResponseHelper(c).InternalServerError(message)
}
