package models

import (
	"time"
)

// Transaction 交易记录模型
type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	CategoryID  uint      `json:"category_id"`
	Amount      float64   `json:"amount"`
	Type        string    `gorm:"size:20" json:"type"` // income: 收入, expense: 支出
	Description string    `gorm:"size:255" json:"description"`
	RecordTime  time.Time `json:"record_time"`
	Location    string    `gorm:"size:255" json:"location"`
	ImageURL    string    `gorm:"size:255" json:"image_url"`
	Tags        string    `gorm:"size:255" json:"tags"` // 以逗号分隔的标签
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	
	// 关联
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category"`
}
