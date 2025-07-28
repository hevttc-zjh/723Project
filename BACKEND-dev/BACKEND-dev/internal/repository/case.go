package repository

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/server"

	"gorm.io/gorm"
)

// CaseRepository 案件仓库
type CaseRepository struct {
	db *gorm.DB
}

// NewCaseRepository 创建案件仓库
func NewCaseRepository() *CaseRepository {
	return &CaseRepository{
		db: server.GetDB(),
	}
}

// GetCaseList 获取案件列表
func (r *CaseRepository) GetCaseList(idCard, name, phone string, page, pageSize int) ([]*model.Case, int64, error) {
	var cases []*model.Case
	var total int64

	// 构建查询
	query := r.db.Model(&model.Case{})

	// 添加查询条件
	if idCard != "" {
		query = query.Where("id_card = ?", idCard)
	}
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if phone != "" {
		query = query.Where("phone = ?", phone)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&cases).Error; err != nil {
		return nil, 0, err
	}

	return cases, total, nil
}

// GetCaseDetail 获取案件详情
func (r *CaseRepository) GetCaseDetail(caseNum string, idCard string) (*model.Case, error) {
	var caseDetail *model.Case
	
	// 构建查询
	query := r.db.Model(&model.Case{})
	
	// 添加查询条件
	query = query.Where("case_num = ? AND id_card = ?", caseNum, idCard)
	
	
	// 执行查询
	if err := query.First(&caseDetail).Error; err != nil {
		return nil, err
	}
	
	return caseDetail, nil
}
