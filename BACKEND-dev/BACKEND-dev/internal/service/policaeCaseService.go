package service

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/repository"
)

// NewPoliceCaseService 创建警情服务
func NewPoliceCaseService() *PoliceCaseService {
	return &PoliceCaseService{
		policeCaseRepo: repository.NewPoliceCaseRepository(),
	}
}

// PoliceCaseService 警情服务
type PoliceCaseService struct {
	policeCaseRepo *repository.PoliceCaseRepository
}

// 获取警情列表
func (s *PoliceCaseService) GetPoliceCaseList(paramsStr string, pageInt int, pageSizeInt int) ([]*model.PoliceCase, int64, error) {
	return s.policeCaseRepo.GetPoliceCaseList(paramsStr, pageInt, pageSizeInt)
}

// 获取警情详情
func GetPoliceCaseDetail(id int) (*model.PoliceCase, error) {
	return repository.NewPoliceCaseRepository().GetPoliceCaseDetail(id)
}
