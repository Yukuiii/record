package models

import (
	"time"
)

// Category 分类模型
type Category struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:50" json:"name"`
	Type      string    `gorm:"size:20" json:"type"` // income: 收入, expense: 支出
	Icon      string    `gorm:"size:50" json:"icon"`
	Color     string    `gorm:"size:20" json:"color"`
	IsDefault bool      `gorm:"default:false" json:"is_default"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
