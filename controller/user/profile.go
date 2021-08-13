package user

import (
	"gateway/model/user"
	"gateway/sql"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"time"
)

// GetProfileController 获取个人信息
func GetProfileController(c *gin.Context) {
	var form user.ProfileGetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	db := sql.UserInfoTable{}
	if err := db.Query(form.UID); err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	data := user.ProfileGetResModel{
		UID:          db.ID,
		Username:     db.Username,
		Birthday:     db.Birthday.Unix(),
		Sex:          db.Sex,
		Avatar:       db.Avatar,
		Motto:        db.Motto,
		Rank:         0,
		SubRank:      0,
		Type:         db.UserType,
		Org:          "四川师范大学",
		Followers:    0,
		Following:    0,
		VideoCounter: 0,
		IsFollowing:  false,
	}
	errno.Perfect(c, &data)
}

// SetProfileController 设置个人信息
func SetProfileController(c *gin.Context) {
	var form user.ProfileSetReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 日期字符串转time.Time
	birthday, err := time.Parse("2006-01-02", form.Birthday)
	if err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 更新数据库
	db := sql.UserInfoTable{
		Username: form.Username,
		Birthday: birthday,
		Sex: form.Sex,
		Motto: form.Motto,
	}
	if err := db.Update(); err != nil {
		errno.Abort(c, errno.TypeMySQLErr)
		return
	}

	errno.Perfect(c, form)
}
