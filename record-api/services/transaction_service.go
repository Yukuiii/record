package services

import (
	"errors"
	"time"

	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/models"
	"github.com/sakura/record-api/repositories"
)

// TransactionService 交易记录服务接口
type TransactionService interface {
	CreateTransaction(userID uint, req api.TransactionRequest) (*models.Transaction, error)
	GetTransactionByID(userID uint, id uint) (*models.Transaction, error)
	GetTransactions(userID uint, params api.TransactionQueryParams) ([]models.Transaction, int64, error)
	UpdateTransaction(userID uint, id uint, req api.TransactionRequest) (*models.Transaction, error)
	DeleteTransaction(userID uint, id uint) error
}

// transactionServiceImpl 交易记录服务实现
type transactionServiceImpl struct {
	transactionRepo repositories.TransactionRepository
	categoryRepo    repositories.CategoryRepository
}

// NewTransactionService 创建交易记录服务
func NewTransactionService() TransactionService {
	return &transactionServiceImpl{
		transactionRepo: repositories.NewTransactionRepository(),
		categoryRepo:    repositories.NewCategoryRepository(),
	}
}

// CreateTransaction 创建交易记录
func (s *transactionServiceImpl) CreateTransaction(userID uint, req api.TransactionRequest) (*models.Transaction, error) {
	// 验证分类是否存在
	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 验证分类类型与交易类型是否匹配
	if category.Type != req.Type {
		return nil, errors.New("分类类型与交易类型不匹配")
	}

	// 设置记录时间，如果没有提供则使用当前时间
	recordTime := time.Now()
	if req.RecordTime != nil {
		recordTime = *req.RecordTime
	}

	// 创建交易记录
	transaction := &models.Transaction{
		UserID:      userID,
		CategoryID:  req.CategoryID,
		Amount:      req.Amount,
		Type:        req.Type,
		Description: req.Description,
		RecordTime:  recordTime,
		Location:    req.Location,
		ImageURL:    req.ImageURL,
		Tags:        req.Tags,
	}

	// 保存交易记录
	if err := s.transactionRepo.Create(transaction); err != nil {
		return nil, err
	}

	// 重新查询以获取关联的分类信息
	return s.transactionRepo.FindByID(transaction.ID)
}

// GetTransactionByID 根据ID获取交易记录
func (s *transactionServiceImpl) GetTransactionByID(userID uint, id uint) (*models.Transaction, error) {
	transaction, err := s.transactionRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, errors.New("交易记录不存在")
	}

	// 验证交易记录是否属于当前用户
	if transaction.UserID != userID {
		return nil, errors.New("无权访问此交易记录")
	}

	return transaction, nil
}

// GetTransactions 获取交易记录列表
func (s *transactionServiceImpl) GetTransactions(userID uint, params api.TransactionQueryParams) ([]models.Transaction, int64, error) {
	// 设置默认分页参数
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 20
	}
	if params.PageSize > 100 {
		params.PageSize = 100
	}

	return s.transactionRepo.FindByUserID(userID, params)
}

// UpdateTransaction 更新交易记录
func (s *transactionServiceImpl) UpdateTransaction(userID uint, id uint, req api.TransactionRequest) (*models.Transaction, error) {
	// 获取交易记录
	transaction, err := s.transactionRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if transaction == nil {
		return nil, errors.New("交易记录不存在")
	}

	// 验证交易记录是否属于当前用户
	if transaction.UserID != userID {
		return nil, errors.New("无权修改此交易记录")
	}

	// 验证分类是否存在
	category, err := s.categoryRepo.FindByID(req.CategoryID)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 验证分类类型与交易类型是否匹配
	if category.Type != req.Type {
		return nil, errors.New("分类类型与交易类型不匹配")
	}

	// 更新交易记录
	transaction.CategoryID = req.CategoryID
	transaction.Amount = req.Amount
	transaction.Type = req.Type
	transaction.Description = req.Description
	if req.RecordTime != nil {
		transaction.RecordTime = *req.RecordTime
	}
	transaction.Location = req.Location
	transaction.ImageURL = req.ImageURL
	transaction.Tags = req.Tags

	// 保存更新
	if err := s.transactionRepo.Update(transaction); err != nil {
		return nil, err
	}

	// 重新查询以获取关联的分类信息
	return s.transactionRepo.FindByID(transaction.ID)
}

// DeleteTransaction 删除交易记录
func (s *transactionServiceImpl) DeleteTransaction(userID uint, id uint) error {
	// 获取交易记录
	transaction, err := s.transactionRepo.FindByID(id)
	if err != nil {
		return err
	}
	if transaction == nil {
		return errors.New("交易记录不存在")
	}

	// 验证交易记录是否属于当前用户
	if transaction.UserID != userID {
		return errors.New("无权删除此交易记录")
	}

	// 删除交易记录
	return s.transactionRepo.Delete(id)
}