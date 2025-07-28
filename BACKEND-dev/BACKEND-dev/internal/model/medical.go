package model

import (
	"time"
)

// Medical 医疗信息模型
type Medical struct {
	ID string `gorm:"primaryKey;column:id;type:uuid"`
	//PersonID              string    `gorm:"column:person_id"`               // 关联Person表的外键
	County                string    `gorm:"column:county"`                   // 县（区）
	Town                  string    `gorm:"column:town"`                     // 乡镇（街道）
	Village               string    `gorm:"column:village"`                  // 村（区）
	Name                  string    `gorm:"column:name"`                     // 姓名
	Gender                string    `gorm:"column:gender"`                   // 性别
	CitizenIdNumber       string    `gorm:"column:citizen_id_number"`        // 公民身份证号
	InsuranceStatus       string    `gorm:"column:insurance_status"`         // 参保状态
	DateOfBirth           time.Time `gorm:"column:date_of_birth"`            // 出生日期
	Age                   int       `gorm:"column:age"`                      // 年龄
	Household             string    `gorm:"column:house_hold"`               // 户籍
	Telephone             string    `gorm:"column:telephone"`                // 手机号
	InsuranceCoverageDate time.Time `gorm:"column:insurance_coverage _date"` // 参保日期
	Backsidate            time.Time `gorm:"column:backsi_date"`              // 退保日期
	PersonnelCategory     string    `gorm:"column:personnel _category"`      // 人员分类
}

// TableName 指定表名
func (Medical) TableName() string {
	return "rsj_temp"
}
