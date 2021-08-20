package sql

import "gorm.io/gorm"

// EventGetTable 消息Get表
type EventGetTable struct {
	gorm.Model
	UID uint
	EID string
}

func (EventGetTable) TableName() string {
	return "event_get"
}

// Create 添加新记录
func (e *EventGetTable) Create() error {
	return db.Create(e).Error
}

// GetRecords 分页获取记录
func (e *EventGetTable) GetRecords(eid string, offset, number int) (total int64, result []EventGetTable, err error) {
	err = db.Model(e).Where("eid = ?", eid).Count(&total).Error
	if err != nil {
		return
	}

	err = db.Where("eid = ?", eid).Offset(offset).Limit(number).Find(&result).Error
	return
}
