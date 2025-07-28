package service

import (
	"fmt"
	// "strconv"
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/repository"
	"risk-insight-system/internal/utils"
	// "strings"
)

type PersonService struct {
	personRepo     *repository.PersonRepository
	medicalService *MedicalService
}

// PersonParams 人员查询参数
type PersonParams struct {
	IdCard string
	Name   string
	Phone  string
	Age    int
	Gender string
}

// GetPersonList 获取人员列表
func (s *PersonService) GetPersonList(paramsStr string, page, pageSize int) ([]*model.Person, int64, error) {
	idCard := utils.ExtractIDCard(paramsStr)
	phone := utils.ExtractMobile(paramsStr)
	name := utils.ExtractName(paramsStr)
	// 只在参数非空时进行校验
	if idCard == "" && utils.ValidateIDCard(idCard) {
		return nil, 0, fmt.Errorf("身份证号格式错误, 请检查输入格式")
	}
	if phone == "" && utils.ValidateMobile(phone) {
		return nil, 0, fmt.Errorf("手机号格式错误, 请检查输入格式")

	}
	if name == "" && utils.ValidateName(name) {
		return nil, 0, fmt.Errorf("姓名格式错误, 请检查输入格式")

	}
	// 确保至少有一个必填参数
	if idCard == "" && phone == "" && name == "" {
		return nil, 0, fmt.Errorf("至少需要提供身份证号、手机号或姓名中的一项,请检查输入格式")
	}
	return s.personRepo.GetPersonList(idCard, name, phone, page, pageSize)
}

// NewPersonService 创建人员服务
func NewPersonService() *PersonService {
	return &PersonService{
		personRepo:     repository.NewPersonRepository(),
		medicalService: NewMedicalService(),
	}
}

// GetPerson
func (s *PersonService) GetPerson(idCard, name, phone string) (*model.PersonSearch, error) {
	idCard = utils.ExtractIDCard(idCard)
	phone = utils.ExtractMobile(phone)
	name = utils.ExtractName(name)
	// 只在参数非空时进行校验
	if idCard == "" && utils.ValidateIDCard(idCard) {
		return nil, fmt.Errorf("身份证号格式错误, 请检查输入格式")
	}
	// if phone == "" && utils.ValidateMobile(phone) {
	// 	return nil, fmt.Errorf("手机号格式错误, 请检查输入格式")
	// }
	// if name == "" && utils.ValidateName(name) {
	// 	return nil,  fmt.Errorf("姓名格式错误, 请检查输入格式")
	// }

	// 确保至少有一个必填参数
	if idCard == "" && phone == "" && name == "" {
		return nil, fmt.Errorf("至少需要提供身份证号、手机号或姓名中的一项,请检查输入格式")
	}

	// 调用repository层获取数据
	persons, err := s.personRepo.GetPerson(idCard, name, phone)
	if err != nil {
		return nil, err
	}
	if len(persons) == 0 {
		return nil, nil
	}
	// 返回包含第一个元素的数组
	return &model.PersonSearch{PersonInfo: []*model.Person{persons[0]}}, nil

}
