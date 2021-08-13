package sql

import (
	"gateway/tools/hashids"
	"gorm.io/gorm"
)

var (
	videoHash = hashids.New("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "abcd", 8)
)

// VideoInfoTable 消息事务信息表
type VideoInfoTable struct {
	gorm.Model
	UID      uint
	VID      string
	Title    string
	Tags     []string
	Cover    string
	Video    string
	Location string
	Status   uint8
}

func (VideoInfoTable) TableName() string {
	return "video_info"
}

// Create 创建一个新消息记录
func (e *VideoInfoTable) Create() error {
	return db.Create(e).Error
}

// AfterCreate 将VID设置为hashids编码形式
func (e *VideoInfoTable) AfterCreate(tx *gorm.DB) error {
	return tx.Model(e).Update("vid", videoHash.Encode(e.ID)).Error
}

// Update 更新消息记录
func (e *VideoInfoTable) Update() error {
	return db.Model(e).Updates(*e).Error
}

// AfterUpdate 更新消息记录 UpdateTime 字段
func (e *VideoInfoTable) AfterUpdate(tx *gorm.DB) error {
	return tx.Model(e).Update("update_time", gorm.Expr("update_time + ?", 1)).Error
}

// Query 查询消息记录
func (e *VideoInfoTable) Query(vid string) error {
	e.VID = vid
	return db.Where(e).Find(e).Error
}

// Delete 删除消息记录
func (e *VideoInfoTable) Delete(vid string) error {
	e.VID = vid
	return db.Where(e).Delete(e).Error
}
