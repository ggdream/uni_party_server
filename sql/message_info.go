package sql

//import (
//	"gateway/tools/hashids"
//	"gorm.io/gorm"
//	"time"
//)
//
//var (
//	eventHash = hashids.New("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "abcd", 8)
//)
//
//// EventInfoTable 消息事务信息表
//type EventInfoTable struct {
//	gorm.Model
//	UID        uint
//	EID        string
//	Type       uint8
//	Title      string
//	Content    string
//	Constraint string
//	UpdateTime uint8
//}
//
//func (EventInfoTable) TableName() string {
//	return "event_info"
//}
//
//// Create 创建一个新消息记录
//func (e *EventInfoTable) Create() error {
//	return db.Create(e).Error
//}
//
//// AfterCreate 将EID设置为hashids编码形式
//func (e *EventInfoTable) AfterCreate(tx *gorm.DB) error {
//	return tx.Model(e).Update("eid", eventHash.Encode(e.ID)).Error
//}
//
//// Update 更新消息记录
//func (e *EventInfoTable) Update() error {
//	return db.Model(e).Updates(*e).Error
//}
//
//// AfterUpdate 更新消息记录 UpdateTime 字段
//func (e *EventInfoTable) AfterUpdate(tx *gorm.DB) error {
//	return tx.Model(e).Update("update_time", gorm.Expr("update_time + ?", 1)).Error
//}
//
//// Query 查询消息记录
//func (e *EventInfoTable) Query(eid string) error {
//	e.EID = eid
//	return db.Where(e).Find(e).Error
//}
//
//// Delete 删除消息记录
//func (e *EventInfoTable) Delete(eid string) error {
//	e.EID = eid
//	return db.Where(e).Delete(e).Error
//}
