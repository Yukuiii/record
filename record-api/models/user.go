package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Email        string    `gorm:"uniqueIndex;size:255" json:"email"`
	Phone        string    `gorm:"uniqueIndex;size:20" json:"phone"`
	Password     string    `gorm:"size:255" json:"-"` // 不返回密码
	Nickname     string    `gorm:"size:50" json:"nickname"`
	Avatar       string    `gorm:"size:255" json:"avatar"`
	Gender       string    `gorm:"size:10" json:"gender"`
	Birthday     time.Time `json:"birthday"`
	RegisterTime time.Time `gorm:"autoCreateTime" json:"register_time"`
	LastLogin    time.Time `json:"last_login"`
	Status       int       `gorm:"default:1" json:"status"` // 1: 正常, 0: 禁用
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeSave 保存前的钩子函数
func (u *User) BeforeSave(tx *gorm.DB) error {
	// 如果密码被修改，则加密密码
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

// ComparePassword 比较密码是否正确
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
