package utils

import (
	"time"
)

// TimeFormat 常用时间格式
const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// GetStartOfDay 获取某天的开始时间（00:00:00）
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取某天的结束时间（23:59:59）
func GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// GetStartOfMonth 获取某月的开始时间
func GetStartOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfMonth 获取某月的结束时间
func GetEndOfMonth(t time.Time) time.Time {
	return GetStartOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// GetStartOfYear 获取某年的开始时间
func GetStartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// GetEndOfYear 获取某年的结束时间
func GetEndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999999999, t.Location())
}

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatDateTime 格式化日期时间
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// ParseDate 解析日期字符串
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse(DateFormat, dateStr)
}

// ParseDateTime 解析日期时间字符串
func ParseDateTime(dateTimeStr string) (time.Time, error) {
	return time.Parse(DateTimeFormat, dateTimeStr)
}

// IsToday 判断是否为今天
func IsToday(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month() && t.Day() == now.Day()
}

// IsThisMonth 判断是否为本月
func IsThisMonth(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year() && t.Month() == now.Month()
}

// IsThisYear 判断是否为今年
func IsThisYear(t time.Time) bool {
	now := time.Now()
	return t.Year() == now.Year()
}

// GetMonthDays 获取某月的天数
func GetMonthDays(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

// GetWeekday 获取星期几（中文）
func GetWeekday(t time.Time) string {
	weekdays := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
	return weekdays[t.Weekday()]
}
