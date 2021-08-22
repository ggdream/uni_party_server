package auth

import (
	"gateway/model/auth"
	"gateway/sql"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
)

// UniversityStructController 获取高校系统的结构
func UniversityStructController(c *gin.Context) {
	var form auth.UniversityStructReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// TODO: 使用一个手机号申请平台，学校证明上传给了平台，待审核状态。审核成功后，分配一个学校官方的用户账号，学校方可上传数据和支付产品
	// 通过账号uid获取到高校名称和高校代码
	universityName := "四川师范大学"
	universityCode := "2222"

	data := make([]sql.StudentRecordTable, len(form.Students))
	for i, student := range form.Students {
		data[i] = sql.StudentRecordTable{
			Name:   student.Name,
			SID:    student.SID,
			CardID: student.CardID,

			Code:       universityCode,
			University: universityName,
			Campus:     student.Campus,
			College:    student.College,
			Grade:      student.Grade,
			Major:      student.Major,
			Class:      student.Class,
		}
	}

	// 学生数据导入至MySQL
	table := sql.StudentRecordTable{}
	if err := table.MultiInsert(data); err != nil {
		errno.New(c, errno.TypeUnknownMistake, nil, "学生数据入库失败")
		return
	}


}

