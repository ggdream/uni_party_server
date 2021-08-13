package sql

import "gorm.io/gorm"

// StudentRecordTable 学生表
type StudentRecordTable struct {
	gorm.Model
	SID        string
	Name       string
	CardID     string
	University string
	Campus     string
	College    string
	Grade      uint8
	Major      string
	Class      string
	Status     uint8
}

func (StudentRecordTable) TableName() string {
	return "student_record"
}

