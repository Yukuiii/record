package services

import (
	"errors"
	"time"

	"github.com/sakura/record-api/api"
	"github.com/sakura/record-api/models"
	"github.com/sakura/record-api/repositories"
	"github.com/sakura/record-api/utils"
)

// AuthService 认证服务接口
type AuthService interface {
	Register(req api.RegisterRequest) (uint, string, error)
	Login(req api.LoginRequest) (*models.User, string, error)
}

// authServiceImpl 认证服务实现
type authServiceImpl struct {
	userRepo repositories.UserRepository
}

// NewAuthService 创建认证服务
func NewAuthService() AuthService {
	return &authServiceImpl{
		userRepo: repositories.NewUserRepository(),
	}
}

// Register 用户注册
func (s *authServiceImpl) Register(req api.RegisterRequest) (uint, string, error) {
	// 检查邮箱是否已存在
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return 0, "", err
	}
	if existingUser != nil {
		return 0, "", errors.New("邮箱已被注册")
	}

	// 检查手机号是否已存在
	existingUser, err = s.userRepo.FindByPhone(req.Phone)
	if err != nil {
		return 0, "", err
	}
	if existingUser != nil {
		return 0, "", errors.New("手机号已被注册")
	}

	// 创建新用户
	newUser := &models.User{
		Email:        req.Email,
		Phone:        req.Phone,
		Password:     req.Password, // 密码在模型的BeforeSave钩子中会被加密
		Nickname:     req.Nickname,
		RegisterTime: time.Now(),
		LastLogin:    time.Now(),
		Status:       1, // 1: 正常
	}

	// 保存用户
	if err := s.userRepo.Create(newUser); err != nil {
		return 0, "", err
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(newUser.ID)
	if err != nil {
		return 0, "", err
	}

	return newUser.ID, token, nil
}

// Login 用户登录
func (s *authServiceImpl) Login(req api.LoginRequest) (*models.User, string, error) {
	// 根据账号(邮箱或手机号)查找用户
	user, err := s.userRepo.FindByEmailOrPhone(req.Account)
	if err != nil {
		return nil, "", err
	}
	if user == nil {
		return nil, "", errors.New("账号不存在")
	}

	// 验证密码
	if !user.ComparePassword(req.Password) {
		return nil, "", errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, "", errors.New("账号已被禁用")
	}

	// 更新最后登录时间
	if err := s.userRepo.UpdateLastLogin(user.ID); err != nil {
		return nil, "", err
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
