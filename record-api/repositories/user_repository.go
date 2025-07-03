package repositories

import (
	"errors"

	"github.com/sakura/record-api/database"
	"github.com/sakura/record-api/models"
	"gorm.io/gorm"
)

// UserRepository 用户数据访问接口
type UserRepository interface {
	Create(user *models.User) error
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	FindByEmailOrPhone(account string) (*models.User, error)
	Update(user *models.User) error
	UpdateLastLogin(id uint) error
}

// userRepositoryImpl 用户数据访问实现
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository 创建用户数据访问实现
func NewUserRepository() UserRepository {
	return &userRepositoryImpl{db: database.DB}
}

// Create 创建用户
func (r *userRepositoryImpl) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByID 根据ID查找用户
func (r *userRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByPhone 根据手机号查找用户
func (r *userRepositoryImpl) FindByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.db.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmailOrPhone 根据邮箱或手机号查找用户
func (r *userRepositoryImpl) FindByEmailOrPhone(account string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ? OR phone = ?", account, account).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update 更新用户信息
func (r *userRepositoryImpl) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// UpdateLastLogin 更新最后登录时间
func (r *userRepositoryImpl) UpdateLastLogin(id uint) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).
		UpdateColumn("last_login", gorm.Expr("NOW()")).Error
}
