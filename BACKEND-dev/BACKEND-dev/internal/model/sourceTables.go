package model

import (
	"time"
)

// SourceTables 数据源表
type SourceTables struct {
	Id         string    `gorm:"primaryKey"`
	Department string    `gorm:"column:department"`
	TableLabel string    `gorm:"column:table_label"`
	TableCode  string    `gorm:"column:table_code"`
	SourceType string    `gorm:"column:source_type"`
	Priority   int       `gorm:"column:priority"`
	Uploader   string    `gorm:"column:uploader"`
	CreateAt   time.Time `gorm:"column:create_at"`
	State      string    `gorm:"column:state"`
	Backup1    string    `gorm:"column:backup1"`
	Backup2    string    `gorm:"column:backup2"`
	Backup3    string    `gorm:"column:backup3"`
}

// TableName 指定表名
func (SourceTables) TableName() string {
	return "source_tables"
}
