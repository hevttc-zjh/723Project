package repository

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/server"

	"gorm.io/gorm"
)

// UserRepository 用户数据库操作
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository 创建用户数据库操作实例
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: server.GetDB(),
	}
}

// GetUserByPoliceId 根据警员ID查询用户
func (r *UserRepository) GetUserByPoliceId(policeId string) (*model.User, error) {
	var user model.User
	err := r.db.Where("police_id = ? AND is_del = 1", policeId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ValidateUser 验证用户登录
func (r *UserRepository) ValidateUser(policeId, password string) (*model.User, error) {
	var user model.User
	err := r.db.Where("police_id = ? AND password = ? AND is_del = 0", policeId, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
