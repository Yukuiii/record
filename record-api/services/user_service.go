package services

import (
	"errors"

	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/models"
	"github.com/sakura/record-api/repositories"
)

// UserService 用户服务接口
type UserService interface {
	GetUserByID(id uint) (*models.User, error)
	UpdateProfile(id uint, req api.UpdateProfileRequest) (*models.User, error)
}

// userServiceImpl 用户服务实现
type userServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewUserService 创建用户服务
func NewUserService() UserService {
	return &userServiceImpl{
		userRepo: repositories.NewUserRepository(),
	}
}

// GetUserByID 根据ID获取用户
func (s *userServiceImpl) GetUserByID(id uint) (*models.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}
	return user, nil
}

// UpdateProfile 更新用户资料
func (s *userServiceImpl) UpdateProfile(id uint, req api.UpdateProfileRequest) (*models.User, error) {
	// 获取用户
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 更新资料
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	
	if req.Birthday != nil {
		user.Birthday = *req.Birthday
	}

	// 保存更新
	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}
