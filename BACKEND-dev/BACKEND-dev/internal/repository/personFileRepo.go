package repository

import (
	"risk-insight-system/internal/model"
	"risk-insight-system/internal/server"
	"strings"
)

// PersonFileRepo 人员档案数据访问层
// 实际上复用personRepo的查询

type PersonFileRepo struct{}

func NewPersonFileRepo() *PersonFileRepo {
	return &PersonFileRepo{}
}

// GetPersonFile 通过身份证号、姓名或手机号查询人员档案
func (r *PersonFileRepo) GetPersonFile(idCard, name, phone string) (*model.PersonFile, error) {
	personRepo := NewPersonRepository()
	// 只要有身份证号、手机号或姓名，查person表返回基本信息
	if idCard != "" || name != "" || phone != "" {
		persons, err := personRepo.GetPerson(idCard, name, phone)
		if err != nil || len(persons) == 0 {
			return nil, err
		}
		return &model.PersonFile{PersonInfo: persons[0]}, nil
	}
	return nil, nil
}
func (r *PersonFileRepo) GetPersonPhoneByIDCard(idCard string) (*model.PersonFile, error) {
	personRepo := NewPersonRepository()
	// 凭借身份证号，查手机号和来源地
	if idCard != "" {
		persons, err := personRepo.GetPerson(idCard, "", "")
		if err != nil || len(persons) == 0 {
			return nil, err
		}
		person := persons[0]
		db := server.GetDB()
		finalPhone := person.Phone
		if finalPhone == "" {
			type PoliceCasePhone struct{ Phone string }
			var policePhones []PoliceCasePhone
			db.Table("police_case").Select("phone").Where("id_card = ?", person.IdCard).Find(&policePhones)
			if len(policePhones) > 0 && policePhones[0].Phone != "" {
				finalPhone = policePhones[0].Phone
			}
		}
		if finalPhone == "" {
			type MedicalPhone struct{ Telephone string }
			var medicalPhones []MedicalPhone
			db.Table("rsj_temp").Select("telephone").Where("citizen_id_number = ?", person.IdCard).Find(&medicalPhones)
			if len(medicalPhones) > 0 && medicalPhones[0].Telephone != "" {
				finalPhone = medicalPhones[0].Telephone
			}
		}
		var department string
		if person.InfoSourceTableName != "" {
			var result struct{ Department string }
			db.Table("table_fields").Select("department").Where("table_name = ?", person.InfoSourceTableName).First(&result)
			department = result.Department
		}
		return &model.PersonFile{
			Phone:      finalPhone,
			Department: department,
		}, nil
	}
	return nil, nil
}

// GetPersonTagsByIDCard 通过身份证号查找标签
func (r *PersonFileRepo) GetPersonTagsByIDCard(idCard string) ([]string, error) {
	if idCard == "" {
		return nil, nil
	}
	personRepo := NewPersonRepository()
	persons, err := personRepo.GetPerson(idCard, "", "")
	if err != nil || len(persons) == 0 {
		return nil, err
	}
	tagStr := persons[0].Tag
	if tagStr == "" {
		return []string{}, nil
	}
	tags := make([]string, 0)
	for _, tag := range strings.Split(tagStr, ",") {
		t := strings.TrimSpace(tag)
		if t != "" {
			tags = append(tags, t)
		}
	}
	return tags, nil
}
