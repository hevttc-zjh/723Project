package service

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/repository"
)

// MedicalService 医疗服务
type MedicalService struct {
	medicalRepository *repository.MedicalRepository
}

// NewMedicalService 创建医疗服务实例
func NewMedicalService() *MedicalService {
	return &MedicalService{
		medicalRepository: repository.NewMedicalRepository(),
	}
}

// GetMedical 获取医疗信息
func (s *MedicalService) GetMedical(idCard, Name, Phone string) ([]*model.Medical, error) {
	// 关联residents表查询医疗信息，关联条件为身份证号
	medicals, err := s.medicalRepository.GetMedicalWithResident(idCard, Name, Phone)
	if err != nil {
		return nil, err
	}

	// 校验年龄是否超过60岁的逻辑可以在这里扩展
	// 遍历医疗记录处理老年医疗信息
	// for _, medical := range medicals {
	// 	if medical.Age > 60 {
	// 		// 处理逻辑
	// 	}
	// }

	return medicals, nil
}
