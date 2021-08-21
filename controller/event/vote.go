package event

import (
	"gateway/middleware"
	"gateway/model/event"
	"gateway/mongo"
	"gateway/tools/errno"
	"github.com/gin-gonic/gin"
	"time"
)

// VoteCreateController 发布投票类消息
func VoteCreateController(c *gin.Context) {
	var form event.VoteCreateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 	消息发布至MongoDB，得到eid
	document := mongo.EventDocument{
		UID:      c.GetUint("uid"),
		Type:     TypeVote,
		Title:    form.Title,
		Content:  form.Content,
		Tags:     form.Tags,
		Datetime: time.Now(),
		Constraint: mongo.VoteField{
			MaxNumber: form.AllowedNumber,
			Deadline:  form.Deadline,
			Options:   form.Options,
		},
	}
	if err := document.Insert(); err != nil {
		errno.Abort(c, errno.TypeEventPublishFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, document)
}

// VoteUpdateController 修改投票类消息
func VoteUpdateController(c *gin.Context) {
	var form event.VoteUpdateReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 对MongoDB进行update操作
	document := mongo.EventDocument{}
	updateDocument := mongo.UpdateEventDocument{
		Title:   form.Title,
		Content: form.Content,
		Tags:    form.Tags,
		Constraint: mongo.VoteField{
			MaxNumber: form.AllowedNumber,
			Deadline:  form.Deadline,
			Options:   form.Options,
		},
	}
	_, err := document.Update(form.EID, &updateDocument)
	if err != nil {
		errno.Abort(c, errno.TypeEventUpdateFailed)
		return
	}

	// 返回结果
	errno.Perfect(c, updateDocument)
}

// VoteDoController 进行投票
func VoteDoController(c *gin.Context) {
	var form event.VoteDoReqModel
	if err := c.ShouldBind(&form); err != nil {
		errno.Abort(c, errno.TypeParamsParsingErr)
		return
	}

	// 查询该消息的细节信息
	e := mongo.EventDocument{}
	eventDocument, err := e.FindOneDetail(form.EID)
	if err != nil {
		errno.Abort(c, errno.TypeMongoErr)
		return
	}
	// 是否为投票消息
	if eventDocument.Type != TypeVote {
		errno.Abort(c, errno.TypeEventTypeErr)
		return
	}

	constraint := eventDocument.Constraint.(mongo.VoteField)
	// 答案个数与选项个数不符
	if len(constraint.Options) != len(form.Answers) {
		errno.Abort(c, errno.TypeEventVoteAOErr)
		return
	}
	// 是否超过截至时间
	if constraint.Deadline < time.Now().Unix() {
		errno.Abort(c, errno.TypeEventDeadlineErr)
		return
	}
	if constraint.MaxNumber < 0 {
		errno.Abort(c, errno.TypeParamsInvalidErr)
		return
	}

	var answerTrue int
	for _, answer := range form.Answers {
		if answer {
			answerTrue++
		}
	}
	// 投票选项数量超过限制
	if answerTrue > constraint.MaxNumber {
		errno.Abort(c, errno.TypeEventTypeErr)
		return
	}

	db := mongo.VoteDocument{
		UID:      c.GetUint(middleware.KeyUID),
		EID:      form.EID,
		Answers:  form.Answers,
		Datetime: time.Now(),
	}
	if err := db.Insert(); err != nil {
		errno.Abort(c, errno.TypeMongoErr)
		return
	}

	data, err := db.Query(form.EID, 0, 0)
	if err != nil {
		errno.Abort(c, errno.TypeMongoErr)
		return
	}

	result := make([]int64, len(form.Answers))
	for _, v := range data {
		for i, answer := range v.Answers {
			if answer {
				result[i] += 1
			}
		}
	}

	ret := &event.VoteDoResModel{
		Total:   len(data),
		Answers: result,
	}
	errno.Perfect(c, ret)
}
