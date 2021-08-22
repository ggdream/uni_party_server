package sql

import "gorm.io/gorm"

// StudentRecordTable 学生表
type StudentRecordTable struct {
	gorm.Model
	UID uint

	Name   string
	SID    string
	CardID string

	Code       string
	University string
	Campus     string
	College    string
	Grade      uint8
	Major      string
	Class      string
}

func (StudentRecordTable) TableName() string {
	return "student_record"
}

// MultiInsert 批量插入多个学生记录
func (s *StudentRecordTable) MultiInsert(data []StudentRecordTable) error {
	return db.Create(&data).Error
}
