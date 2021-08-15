package sql

import (
	"gateway/tools/hashids"
	"gorm.io/gorm"
)

var (
	videoHash = hashids.New("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", "abcd", 8)
)

// VideoInfoTable 视频事务信息表
type VideoInfoTable struct {
	gorm.Model
	UID      uint
	VID      string
	Title    string
	Tags     []string
	Cover    string
	Video    string
	Location string
	Status   uint8	// 待审核、未通过、正常、锁定、被平台删除、用户自行删除
}

func (VideoInfoTable) TableName() string {
	return "video_info"
}

// Create 创建一个新视频记录
func (e *VideoInfoTable) Create() error {
	return db.Create(e).Error
}

// AfterCreate 将VID设置为hashids编码形式
func (e *VideoInfoTable) AfterCreate(tx *gorm.DB) error {
	return tx.Model(e).Update("vid", videoHash.Encode(e.ID)).Error
}

// Update 更新视频记录
func (e *VideoInfoTable) Update() error {
	return db.Model(e).Updates(*e).Error
}

// AfterUpdate 更新视频记录 UpdateTime 字段
func (e *VideoInfoTable) AfterUpdate(tx *gorm.DB) error {
	return tx.Model(e).Update("update_time", gorm.Expr("update_time + ?", 1)).Error
}

// Query 查询视频记录
func (e *VideoInfoTable) Query(vid string) error {
	e.VID = vid
	return db.Where(e).Find(e).Error
}

// QueryPage 分页查询视频记录
func (e *VideoInfoTable) QueryPage(uid uint, offset, number int) ([]VideoInfoTable, error) {
	var result []VideoInfoTable
	e.UID = uid

	err := db.Where(e).Limit(number).Offset(offset).Order("id DESC").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// QueryIn 多值查询视频记录
func (e *VideoInfoTable) QueryIn(uid uint, vidList []string) ([]VideoInfoTable, error) {
	var result []VideoInfoTable
	e.UID = uid

	err := db.Where("vid IN ?", vidList).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Delete 删除视频记录
func (e *VideoInfoTable) Delete(vid string) error {
	e.VID = vid
	return db.Where(e).Delete(e).Error
}
