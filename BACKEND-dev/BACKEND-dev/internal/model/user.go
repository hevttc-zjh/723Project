package model

import (
	"time"

	
)

// User 用戶模型
type User struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	PoliceId   string         `json:"police_id" gorm:"uniqueIndex;not null;comment:警員ID"`
	Password   string         `json:"password" gorm:"not null;comment:密碼"`
	Role       string         `json:"role" gorm:"comment:角色"`
	CreateTime  time.Time      `json:"create_time"`
	UpdateTime  time.Time      `json:"update_time"`
	IsDel       int            `json:"is_del" gorm:"default:0;comment:是否删除 0:未删除 1:已删除"`
	IdCard      string         `json:"id_card" gorm:"comment:身份证号"`
	NickName    string         `json:"nick_name" gorm:"comment:姓名"`
	Apartment   string         `json:"apartment" gorm:"comment:部门"`
	Phone       string         `json:"phone" gorm:"comment:手机号"`
	UpdateBy    string         `json:"update_by" gorm:"comment:更新人"`
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}
