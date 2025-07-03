package services

import (
	"errors"

	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/models"
	"github.com/sakura/record-api/repositories"
)

// CategoryService 分类服务接口
type CategoryService interface {
	GetAllCategories() ([]models.Category, error)
	GetCategoriesByType(categoryType string) ([]models.Category, error)
	CreateCategory(req api.CategoryRequest) (*models.Category, error)
	UpdateCategory(id uint, req api.CategoryRequest) (*models.Category, error)
	DeleteCategory(id uint) error
	GetCategoryByID(id uint) (*models.Category, error)
}

// categoryServiceImpl 分类服务实现
type categoryServiceImpl struct {
	categoryRepo repositories.CategoryRepository
}

// NewCategoryService 创建分类服务
func NewCategoryService() CategoryService {
	return &categoryServiceImpl{
		categoryRepo: repositories.NewCategoryRepository(),
	}
}

// GetAllCategories 获取所有分类
func (s *categoryServiceImpl) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepo.FindAll()
}

// GetCategoriesByType 根据类型获取分类
func (s *categoryServiceImpl) GetCategoriesByType(categoryType string) ([]models.Category, error) {
	// 验证分类类型
	if categoryType != "income" && categoryType != "expense" {
		return nil, errors.New("无效的分类类型")
	}
	return s.categoryRepo.FindByType(categoryType)
}

// CreateCategory 创建分类
func (s *categoryServiceImpl) CreateCategory(req api.CategoryRequest) (*models.Category, error) {
	// 检查分类名称是否已存在
	existingCategory, err := s.categoryRepo.FindByName(req.Name)
	if err != nil {
		return nil, err
	}
	if existingCategory != nil {
		return nil, errors.New("分类名称已存在")
	}

	// 创建新分类
	newCategory := &models.Category{
		Name:      req.Name,
		Type:      req.Type,
		Icon:      req.Icon,
		Color:     req.Color,
		IsDefault: false, // 用户创建的分类不是默认分类
	}

	// 保存分类
	if err := s.categoryRepo.Create(newCategory); err != nil {
		return nil, err
	}

	return newCategory, nil
}

// UpdateCategory 更新分类
func (s *categoryServiceImpl) UpdateCategory(id uint, req api.CategoryRequest) (*models.Category, error) {
	// 获取分类
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}

	// 检查是否为默认分类
	if category.IsDefault {
		return nil, errors.New("默认分类不能修改")
	}

	// 检查新名称是否与其他分类重复
	if req.Name != category.Name {
		existingCategory, err := s.categoryRepo.FindByName(req.Name)
		if err != nil {
			return nil, err
		}
		if existingCategory != nil && existingCategory.ID != id {
			return nil, errors.New("分类名称已存在")
		}
	}

	// 更新分类信息
	category.Name = req.Name
	category.Type = req.Type
	category.Icon = req.Icon
	category.Color = req.Color

	// 保存更新
	if err := s.categoryRepo.Update(category); err != nil {
		return nil, err
	}

	return category, nil
}

// DeleteCategory 删除分类
func (s *categoryServiceImpl) DeleteCategory(id uint) error {
	// 获取分类
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return err
	}
	if category == nil {
		return errors.New("分类不存在")
	}

	// 检查是否为默认分类
	if category.IsDefault {
		return errors.New("默认分类不能删除")
	}

	// 删除分类
	return s.categoryRepo.Delete(id)
}

// GetCategoryByID 根据ID获取分类
func (s *categoryServiceImpl) GetCategoryByID(id uint) (*models.Category, error) {
	category, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, errors.New("分类不存在")
	}
	return category, nil
}