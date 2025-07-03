package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword 密码加密
func HashPassword(password string) (string, error) {
	// 使用默认成本加密密码
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword 比较密码是否匹配
func CheckPassword(password, hash string) bool {
	// 比较密码和哈希值
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ValidatePassword 验证密码强度
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("密码长度至少6位")
	}
	if len(password) > 128 {
		return errors.New("密码长度不能超过128位")
	}

	// 检查是否包含至少一个字母和一个数字
	hasLetter := false
	hasNumber := false

	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		}
		if unicode.IsNumber(char) {
			hasNumber = true
		}
	}

	if !hasLetter || !hasNumber {
		return errors.New("密码必须包含至少一个字母和一个数字")
	}

	return nil
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

// ValidateEmail 验证邮箱格式
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePhone 验证手机号格式（中国大陆）
func ValidatePhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phone)
}
