package services

import (
	"errors"

	"github.com/sakura/record-api/repositories"
)

// StatisticsService 统计分析服务接口
type StatisticsService interface {
	GetMonthlyStatistics(userID uint, year int, month int) (map[string]interface{}, error)
	GetYearlyStatistics(userID uint, year int) (map[string]interface{}, error)
}

// statisticsServiceImpl 统计分析服务实现
type statisticsServiceImpl struct {
	transactionRepo repositories.TransactionRepository
}

// NewStatisticsService 创建统计分析服务
func NewStatisticsService() StatisticsService {
	return &statisticsServiceImpl{
		transactionRepo: repositories.NewTransactionRepository(),
	}
}

// GetMonthlyStatistics 获取月度统计
func (s *statisticsServiceImpl) GetMonthlyStatistics(userID uint, year int, month int) (map[string]interface{}, error) {
	// 验证参数
	if year < 1900 || year > 2100 {
		return nil, errors.New("无效的年份")
	}
	if month < 1 || month > 12 {
		return nil, errors.New("无效的月份")
	}

	// 调用仓储层获取统计数据
	return s.transactionRepo.GetMonthlyStatistics(userID, year, month)
}

// GetYearlyStatistics 获取年度统计
func (s *statisticsServiceImpl) GetYearlyStatistics(userID uint, year int) (map[string]interface{}, error) {
	// 验证参数
	if year < 1900 || year > 2100 {
		return nil, errors.New("无效的年份")
	}

	// 调用仓储层获取统计数据
	return s.transactionRepo.GetYearlyStatistics(userID, year)
}