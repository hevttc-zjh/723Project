package model

import (
	"time"
)

// person 人员模型
type Person struct {
	ID                  string    `gorm:"primaryKey;column:id"`
	IdCard              string    `gorm:"unique;column:id_card"`
	Name                string    `gorm:"column:name"`
	Gender              string    `gorm:"column:gender"`
	BrithDate           string    `gorm:"column:birth_date"`
	Address             string    `gorm:"column:address"`
	InfoSourceTableName string    `gorm:"column:info_source_table_name"`
	InfoSourcePriority  int       `gorm:"column:info_source_priority"`
	CreateAt            time.Time `gorm:"column:create_at"`
	Phone               string    `gorm:"column:phone"`
	State               string    `gorm:"column:state"`
	Tag                 string    `gorm:"column:tag"`
}

// TableName 指定表名
func (Person) TableName() string {
	return "residents"
}
