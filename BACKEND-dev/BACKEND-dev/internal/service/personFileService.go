package service

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/repository"
	//"risk-insight-system/internal/service"
)

// PersonFileService 人员档案服务层

type PersonFileService struct {
	repo *repository.PersonFileRepo
}

func NewPersonFileService() *PersonFileService {
	return &PersonFileService{
		repo: repository.NewPersonFileRepo(),
	}
}

// GetPersonFile 通过身份证号、姓名或手机号查询人员档案
func (s *PersonFileService) GetPersonFile(idCard, name, phone string) (*model.PersonFile, error) {
	return s.repo.GetPersonFile(idCard, name, phone)
}

func (s *PersonFileService) GetPersonPhoneByIDCard(idCard string) (*model.PersonFile, error) {
	return s.repo.GetPersonPhoneByIDCard(idCard)
}

func (s *PersonFileService) GetPersonTagsByIDCard(idCard string) ([]string, error) {
	return s.repo.GetPersonTagsByIDCard(idCard)
}

func (s *PersonFileService) GetPoliceCaseList(paramsStr string, pageInt, pageSizeInt int) ([]*model.PoliceCase, int64, error) {
	policeCaseService := NewPoliceCaseService()
	return policeCaseService.GetPoliceCaseList(paramsStr, pageInt, pageSizeInt)
}

func (s *PersonFileService) GetCaseList(paramsStr string, pageInt, pageSizeInt int) ([]*model.Case, int64, error) {
	caseService := NewCaseService()
	return caseService.GetCaseList(paramsStr, pageInt, pageSizeInt)
}

func (s *PersonFileService) GetPersonWithMedicalByIDCard(idCard string) (*model.PersonFile, error) {
	personRepo := repository.NewPersonRepository()
	persons, err := personRepo.GetPerson(idCard, "", "")
	if err != nil || len(persons) == 0 {
		return nil, err
	}
	medicalRepo := repository.NewMedicalRepository()
	medicals, err := medicalRepo.GetMedicalWithResident(idCard, "", "")
	if err != nil {
		return nil, err
	}
	return &model.PersonFile{
		PersonInfo:  persons[0],
		MedicalInfo: medicals,
	}, nil
}
