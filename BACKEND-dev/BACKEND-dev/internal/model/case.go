package model

import (
	"time"
)

// case 案件模型
type Case struct {
	// 案件编号
	caseNum string `gorm:column:case_num`
	// 案件类型
	caseType string `gorm:cloumn:case_type`
	// 案件名称
	caseName string `gorm:cloumn:case_name`
	// 案件类别
	caseSort string `gorm:cloumn:case_sort`
	// 姓名
	name     string  `gorm:cloumn:name`
	// 民族
	nation string    `gorm:cloumn:nation`
	// 性别
	gender string    `gorm:cloumn:gender`
	// 出生日期
	birthdate time.Time `gorm:cloumn:birthdate`
	// 年龄
	age int `gorm:cloumn:age`
	// 政治面貌
	politicsStatus string `gorm:cloumn:politics_status`
	// 现住址
	presentAddress string `gorm:cloumn:present_address`
	// 手机号
	phone string `gorm:cloumn:phone`
	// 家属电话
	familyPhone string `gorm:cloumn:family_phone`
	// 身份证号
	idCard string `gorm:cloumn:id_card`
	// 办案单位
	caseHandlingUnit string `gorm:cloumn:case_handling_unit`
	// 人员登记日期
	personnelRegistrationDate time.Time `gorm:cloumn:personnel_registration_date`
	// 办案人员
	personnelHandlingCase string `gorm:cloumn:personnel_handling_case`
	
}