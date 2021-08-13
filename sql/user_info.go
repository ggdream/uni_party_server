package sql

import (
	"gorm.io/gorm"
	"time"
)

// UserInfoTable 用户信息表
type UserInfoTable struct {
	gorm.Model
	IdentifyType uint8
	IdentifyNo   string
	Telephone    string
	EMail        string
	Motto        string
	Birthday     time.Time
	Sex          string
	Avatar       string
	UserType     uint8
	Username     string
	Password     string
	Salt         string
	Status       uint8
}

func (UserInfoTable) TableName() string {
	return "user_info"
}

// Create 新增用户
func (u *UserInfoTable) Create() error {
	return db.Create(u).Error
}

// Update 更新用户信息
func (u *UserInfoTable) Update() error {
	return db.Model(u).Updates(*u).Error
}

// UpdateStatus 更新用户状态
func (u *UserInfoTable) UpdateStatus(status uint8) error {
	return db.Model(u).Update("status", status).Error
}

// Query 查询用户信息
func (u *UserInfoTable) Query(uid uint) error {
	return db.Find(u, uid).Error
}

// QueryOnly 只查询用户的手机号、邮箱、密码、盐值
func (u *UserInfoTable) QueryOnly(uid uint) error {
	return db.Select([]string{"telephone", "email", "password", "salt"}).Find(u, uid).Error
}

// Delete 删除用户
func (u *UserInfoTable) Delete(uid uint) error {
	return db.Delete(u, uid).Error
}

// QueryUserByPhone 通过手机号查找用户
func (u *UserInfoTable) QueryUserByPhone(telephone string) error {
	return db.Where("telephone = ?", telephone).Find(u).Error
}

// QueryUserByEMail 通过邮箱查找用户
func (u *UserInfoTable) QueryUserByEMail(email string) error {
	return db.Where("email = ?", email).Find(u).Error
}
