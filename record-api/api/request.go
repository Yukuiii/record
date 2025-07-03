package api

import "time"

// 用户相关请求

// RegisterRequest 用户注册请求
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Nickname string `json:"nickname" binding:"required"`
}

// LoginRequest 用户登录请求
type LoginRequest struct {
	Account  string `json:"account" binding:"required"` // 邮箱或手机号
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest 更新用户信息请求
type UpdateProfileRequest struct {
	Nickname string     `json:"nickname"`
	Avatar   string     `json:"avatar"`
	Gender   string     `json:"gender" binding:"omitempty,oneof=male female other"`
	Birthday *time.Time `json:"birthday"`
}

// 分类相关请求

// CategoryRequest 分类请求
type CategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=income expense"`
	Icon  string `json:"icon" binding:"required"`
	Color string `json:"color" binding:"required"`
}

// 交易记录相关请求

// TransactionRequest 交易记录请求
type TransactionRequest struct {
	CategoryID  uint       `json:"category_id" binding:"required"`
	Amount      float64    `json:"amount" binding:"required,gt=0"`
	Type        string     `json:"type" binding:"required,oneof=income expense"`
	Description string     `json:"description"`
	RecordTime  *time.Time `json:"record_time" binding:"required"`
	Location    string     `json:"location"`
	ImageURL    string     `json:"image_url"`
	Tags        string     `json:"tags"`
}

// TransactionQueryParams 交易记录查询参数
type TransactionQueryParams struct {
	Type       string `form:"type" binding:"omitempty,oneof=income expense"`
	StartDate  string `form:"start_date"`
	EndDate    string `form:"end_date"`
	CategoryID uint   `form:"category_id"`
	Page       int    `form:"page,default=1" binding:"min=1"`
	PageSize   int    `form:"page_size,default=20" binding:"min=1,max=100"`
}

// 统计相关请求

// MonthlyStatisticsQueryParams 月度统计查询参数
type MonthlyStatisticsQueryParams struct {
	Year  int `form:"year" binding:"required"`
	Month int `form:"month" binding:"required,min=1,max=12"`
}

// YearlyStatisticsQueryParams 年度统计查询参数
type YearlyStatisticsQueryParams struct {
	Year int `form:"year" binding:"required"`
}
