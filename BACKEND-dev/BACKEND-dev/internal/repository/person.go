package repository

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/server"

	"gorm.io/gorm"
)

// PersonRepository 人员仓库操作
type PersonRepository struct {
	db *gorm.DB
}

// GetPersonList 分页获取人员列表
func (r *PersonRepository) GetPersonList(idCard, name, phone string, page, pageSize int) ([]*model.Person, int64, error) {
	var persons []*model.Person
	var total int64

	// 构建查询
	query := r.db.Model(&model.Person{})

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
	if err := query.Offset(offset).Limit(pageSize).Find(&persons).Error; err != nil {
		return nil, 0, err
	}

	return persons, total, nil
}

// NewPersonRepository 创建人员仓库操作实例
func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		db: server.GetDB(),
	}
}

// GetPerson 获取人员信息
func (r *PersonRepository) GetPerson(idCard, name, phone string) ([]*model.Person, error) {
	var person []*model.Person
	if err := r.db.Where("id_card = ? OR name = ? OR phone = ?", idCard, name, phone).Find(&person).Error; err != nil {
		return nil, err
	}
	return person, nil
}
