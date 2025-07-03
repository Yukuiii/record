package repositories

import (
	"errors"
	"time"

	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/database"
	"github.com/sakura/record-api/models"
	"gorm.io/gorm"
)

// TransactionRepository 交易记录数据访问接口
type TransactionRepository interface {
	Create(transaction *models.Transaction) error
	FindByID(id uint) (*models.Transaction, error)
	FindByUserID(userID uint, params api.TransactionQueryParams) ([]models.Transaction, int64, error)
	Update(transaction *models.Transaction) error
	Delete(id uint) error
	GetMonthlyStatistics(userID uint, year int, month int) (map[string]interface{}, error)
	GetYearlyStatistics(userID uint, year int) (map[string]interface{}, error)
}

// transactionRepositoryImpl 交易记录数据访问实现
type transactionRepositoryImpl struct {
	db *gorm.DB
}

// NewTransactionRepository 创建交易记录数据访问实现
func NewTransactionRepository() TransactionRepository {
	return &transactionRepositoryImpl{db: database.DB}
}

// Create 创建交易记录
func (r *transactionRepositoryImpl) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

// FindByID 根据ID查找交易记录
func (r *transactionRepositoryImpl) FindByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Category").First(&transaction, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &transaction, nil
}

// FindByUserID 根据用户ID查找交易记录（支持分页和筛选）
func (r *transactionRepositoryImpl) FindByUserID(userID uint, params api.TransactionQueryParams) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var total int64

	// 构建查询
	query := r.db.Model(&models.Transaction{}).Where("user_id = ?", userID)

	// 添加筛选条件
	if params.Type != "" {
		query = query.Where("type = ?", params.Type)
	}

	if params.CategoryID != 0 {
		query = query.Where("category_id = ?", params.CategoryID)
	}

	if params.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", params.StartDate)
		if err == nil {
			query = query.Where("record_time >= ?", startDate)
		}
	}

	if params.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", params.EndDate)
		if err == nil {
			// 结束日期包含当天，所以加一天
			endDate = endDate.AddDate(0, 0, 1)
			query = query.Where("record_time < ?", endDate)
		}
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (params.Page - 1) * params.PageSize
	err := query.Preload("Category").
		Order("record_time DESC, created_at DESC").
		Offset(offset).
		Limit(params.PageSize).
		Find(&transactions).Error

	return transactions, total, err
}

// Update 更新交易记录
func (r *transactionRepositoryImpl) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

// Delete 删除交易记录
func (r *transactionRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Transaction{}, id).Error
}

// GetMonthlyStatistics 获取月度统计
func (r *transactionRepositoryImpl) GetMonthlyStatistics(userID uint, year int, month int) (map[string]interface{}, error) {
	// 构建时间范围
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, 0)

	// 查询收入统计
	var incomeTotal float64
	var incomeCount int64
	err := r.db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ? AND record_time >= ? AND record_time < ?", userID, "income", startDate, endDate).
		Select("COALESCE(SUM(amount), 0) as total, COUNT(*) as count").
		Row().Scan(&incomeTotal, &incomeCount)
	if err != nil {
		return nil, err
	}

	// 查询支出统计
	var expenseTotal float64
	var expenseCount int64
	err = r.db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ? AND record_time >= ? AND record_time < ?", userID, "expense", startDate, endDate).
		Select("COALESCE(SUM(amount), 0) as total, COUNT(*) as count").
		Row().Scan(&expenseTotal, &expenseCount)
	if err != nil {
		return nil, err
	}

	// 查询分类统计
	var categoryStats []map[string]interface{}
	err = r.db.Model(&models.Transaction{}).
		Select("categories.name as category_name, categories.type as category_type, COALESCE(SUM(transactions.amount), 0) as total, COUNT(*) as count").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.id").
		Where("transactions.user_id = ? AND transactions.record_time >= ? AND transactions.record_time < ?", userID, startDate, endDate).
		Group("categories.id, categories.name, categories.type").
		Having("COUNT(*) > 0").
		Scan(&categoryStats).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"year":  year,
		"month": month,
		"income": map[string]interface{}{
			"total": incomeTotal,
			"count": incomeCount,
		},
		"expense": map[string]interface{}{
			"total": expenseTotal,
			"count": expenseCount,
		},
		"balance":          incomeTotal - expenseTotal,
		"category_stats":   categoryStats,
		"total_count":      incomeCount + expenseCount,
	}, nil
}

// GetYearlyStatistics 获取年度统计
func (r *transactionRepositoryImpl) GetYearlyStatistics(userID uint, year int) (map[string]interface{}, error) {
	// 构建时间范围
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(1, 0, 0)

	// 查询年度收入统计
	var incomeTotal float64
	var incomeCount int64
	err := r.db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ? AND record_time >= ? AND record_time < ?", userID, "income", startDate, endDate).
		Select("COALESCE(SUM(amount), 0) as total, COUNT(*) as count").
		Row().Scan(&incomeTotal, &incomeCount)
	if err != nil {
		return nil, err
	}

	// 查询年度支出统计
	var expenseTotal float64
	var expenseCount int64
	err = r.db.Model(&models.Transaction{}).
		Where("user_id = ? AND type = ? AND record_time >= ? AND record_time < ?", userID, "expense", startDate, endDate).
		Select("COALESCE(SUM(amount), 0) as total, COUNT(*) as count").
		Row().Scan(&expenseTotal, &expenseCount)
	if err != nil {
		return nil, err
	}

	// 查询月度趋势
	var monthlyTrends []map[string]interface{}
	for month := 1; month <= 12; month++ {
		monthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		monthEnd := monthStart.AddDate(0, 1, 0)

		var monthIncome, monthExpense float64
		r.db.Model(&models.Transaction{}).
			Where("user_id = ? AND type = ? AND record_time >= ? AND record_time < ?", userID, "income", monthStart, monthEnd).
			Select("COALESCE(SUM(amount), 0)").Row().Scan(&monthIncome)

		r.db.Model(&models.Transaction{}).
			Where("user_id = ? AND type = ? AND record_time >= ? AND record_time < ?", userID, "expense", monthStart, monthEnd).
			Select("COALESCE(SUM(amount), 0)").Row().Scan(&monthExpense)

		monthlyTrends = append(monthlyTrends, map[string]interface{}{
			"month":   month,
			"income":  monthIncome,
			"expense": monthExpense,
			"balance": monthIncome - monthExpense,
		})
	}

	// 查询分类统计
	var categoryStats []map[string]interface{}
	err = r.db.Model(&models.Transaction{}).
		Select("categories.name as category_name, categories.type as category_type, COALESCE(SUM(transactions.amount), 0) as total, COUNT(*) as count").
		Joins("LEFT JOIN categories ON transactions.category_id = categories.id").
		Where("transactions.user_id = ? AND transactions.record_time >= ? AND transactions.record_time < ?", userID, startDate, endDate).
		Group("categories.id, categories.name, categories.type").
		Having("COUNT(*) > 0").
		Scan(&categoryStats).Error
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"year": year,
		"income": map[string]interface{}{
			"total": incomeTotal,
			"count": incomeCount,
		},
		"expense": map[string]interface{}{
			"total": expenseTotal,
			"count": expenseCount,
		},
		"balance":          incomeTotal - expenseTotal,
		"monthly_trends":   monthlyTrends,
		"category_stats":   categoryStats,
		"total_count":      incomeCount + expenseCount,
	}, nil
}