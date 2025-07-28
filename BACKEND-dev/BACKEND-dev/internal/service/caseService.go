package service

import (
	"fmt"
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/repository"
	"risk-insight-system/internal/utils"
)

// GetCaseDetail 获取案件详情
func (s *CaseService) GetCaseDetail(caseNum string, idCard string) (*model.Case, error) {
	if caseNum == "" {
		return nil, fmt.Errorf("案件编号不能为空")
	}
	
	// 身份证号非空时进行校验
	if idCard != "" && utils.ValidateIDCard(idCard) {
		return nil, fmt.Errorf("身份证号格式错误")
	}
	
	return s.caseRepo.GetCaseDetail(caseNum, idCard)
}

type CaseService struct {
	caseRepo *repository.CaseRepository
}

// NewCaseService 创建案件服务
func NewCaseService() *CaseService {
	return &CaseService{
		caseRepo: repository.NewCaseRepository(),
	}
}

// 案件查询参数CaseParams
type CaseParams struct {
	idCard string
	name   string
	phone  string
}

// GetCaseList 获取案件列表
func (s *CaseService) GetCaseList(paramsStr string, page, pageSize int) ([]*model.Case, int64, error) {
	idCard := utils.ExtractIDCard(paramsStr)
	phone := utils.ExtractMobile(paramsStr)
	name := utils.ExtractName(paramsStr)
	// 参数非空时校验
	if idCard == "" && utils.ValidateIDCard(idCard) {
		return nil, 0, fmt.Errorf("身份证号格式错误,请检查输入格式")
	}
	if phone == "" && utils.ValidateMobile(phone) {
		return nil, 0, fmt.Errorf("手机号格式错误,请检查输入格式")
	}
	if name == "" && utils.ValidateName(name) {
		return nil, 0, fmt.Errorf("姓名格式错误,请检查输入格式")
	}
	// 校验必填参数存在性（示例：至少需要姓名或手机号）
	if name == "" && phone == "" && idCard == "" {
		return nil, 0, fmt.Errorf("缺少有效查询条件，需要提供姓名、手机号或身份证号,请检查输入格式")
	}
	// 调用repository层获取数据
	return s.caseRepo.GetCaseList(idCard, name, phone, page, pageSize)

}
