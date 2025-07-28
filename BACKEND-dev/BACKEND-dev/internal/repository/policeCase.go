package repository

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/server"

	"gorm.io/gorm"
)

// PoliceCaseRepository 警情仓库
type PoliceCaseRepository struct {
	db *gorm.DB
}

// GetPoliceCaseList 获取警情列表
func (r *PoliceCaseRepository) GetPoliceCaseList(paramsStr string, pageInt int, pageSizeInt int) ([]*model.PoliceCase, int64, error) {
	var policeCases []*model.PoliceCase
	err := r.db.Find(&policeCases).Error
	if err != nil {
		return nil, 0, err
	}
	return policeCases, 0, nil
}

// GetPoliceCaseDetail 获取警情详情
func (r *PoliceCaseRepository) GetPoliceCaseDetail(id int) (*model.PoliceCase, error) {
	var policeCase model.PoliceCase
	err := r.db.First(&policeCase, id).Error
	if err != nil {
		return nil, err
	}
	return &policeCase, nil
}

// NewPoliceCaseRepository 创建警情仓库实例
func NewPoliceCaseRepository() *PoliceCaseRepository {
	return &PoliceCaseRepository{
		db: server.DB,
	}
}
