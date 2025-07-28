package repository

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/server"

	"gorm.io/gorm"
)

// MedicalRepository 医疗数据仓库
type MedicalRepository struct {
	db *gorm.DB
}

// NewMedicalRepository 创建医疗数据仓库实例
func NewMedicalRepository() *MedicalRepository {
	return &MedicalRepository{
		db: server.DB,
	}
}

// GetMedicalWithResident 获取医疗数据
func (r *MedicalRepository) GetMedicalWithResident(idCard, name, phone string) ([]*model.Medical, error) {
	var medicals []*model.Medical
	err := r.db.Table("rsj_temp").Joins("JOIN residents ON rsj_temp.citizen_id_number = residents.id_card").
		Where("rsj_temp.citizen_id_number = ? OR rsj_temp.name = ? OR rsj_temp.telephone = ?", idCard, name, phone).
		Find(&medicals).Error
	if err != nil {
		return nil, err
	}
	return medicals, nil
	// var medical model.Medical
	// err := r.db.Where("id_card = ? OR name = ? OR phone = ?", idCard, name, phone).First(&medical).Error
	// if err != nil {
	// 	return nil, err
	// }
	// return &medical, nil
}
