package api

import "time"

// Response API标准响应格式
type Response struct {
	Code    int         `json:"code"`    // 状态码: 200成功, 400请求错误, 401未授权, 403禁止访问, 404未找到, 500服务器错误
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
	Timestamp int64       `json:"timestamp"` // 响应时间戳
}

// NewResponse 创建标准响应
func NewResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
		Timestamp: time.Now().Unix(),
	}
}

// SuccessResponse 创建成功响应
func SuccessResponse(data interface{}) Response {
	return NewResponse(200, "success", data)
}

// ErrorResponse 创建错误响应
func ErrorResponse(code int, message string) Response {
	return NewResponse(code, message, nil)
}
