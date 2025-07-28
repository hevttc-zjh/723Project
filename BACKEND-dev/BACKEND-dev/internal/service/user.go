package service

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/repository"
)

// UserService 用户业务逻辑
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService 创建用户业务逻辑实例
func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(),
	}
}

// Login 用户登录
func (s *UserService) Login(policeId, password string) (*model.User, error) {
	// 验证用户
	user, err := s.userRepo.ValidateUser(policeId, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByPoliceId 根据警员ID获取用户信息
func (s *UserService) GetUserByPoliceId(policeId string) (*model.User, error) {
	return s.userRepo.GetUserByPoliceId(policeId)
}
