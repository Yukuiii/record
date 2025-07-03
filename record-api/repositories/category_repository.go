package repositories

import (
	"errors"

	"github.com/sakura/record-api/database"
	"github.com/sakura/record-api/models"
	"gorm.io/gorm"
)

// CategoryRepository 分类数据访问接口
type CategoryRepository interface {
	Create(category *models.Category) error
	FindAll() ([]models.Category, error)
	FindByID(id uint) (*models.Category, error)
	FindByType(categoryType string) ([]models.Category, error)
	Update(category *models.Category) error
	Delete(id uint) error
	FindByName(name string) (*models.Category, error)
}

// categoryRepositoryImpl 分类数据访问实现
type categoryRepositoryImpl struct {
	db *gorm.DB
}

// NewCategoryRepository 创建分类数据访问实现
func NewCategoryRepository() CategoryRepository {
	return &categoryRepositoryImpl{db: database.DB}
}

// Create 创建分类
func (r *categoryRepositoryImpl) Create(category *models.Category) error {
	return r.db.Create(category).Error
}

// FindAll 获取所有分类
func (r *categoryRepositoryImpl) FindAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Order("type ASC, created_at ASC").Find(&categories).Error
	return categories, err
}

// FindByID 根据ID查找分类
func (r *categoryRepositoryImpl) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

// FindByType 根据类型查找分类
func (r *categoryRepositoryImpl) FindByType(categoryType string) ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Where("type = ?", categoryType).Order("created_at ASC").Find(&categories).Error
	return categories, err
}

// Update 更新分类
func (r *categoryRepositoryImpl) Update(category *models.Category) error {
	return r.db.Save(category).Error
}

// Delete 删除分类
func (r *categoryRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

// FindByName 根据名称查找分类
func (r *categoryRepositoryImpl) FindByName(name string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("name = ?", name).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}