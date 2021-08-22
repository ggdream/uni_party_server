package mongo

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	eventCollectionName = "event"
)

// EventDocument 消息文档体
type EventDocument struct {
	UID      uint   `json:"uid,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Username string `json:"username,omitempty"`

	EID      string    `json:"eid,omitempty"`
	Type     uint8     `json:"type,omitempty"`
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	Datetime time.Time `json:"datetime"`
	// 不同消息对应不同的约束
	Constraint interface{} `json:"constraint,omitempty"`
}

// SimpleEventDocument 简略消息文档体
type SimpleEventDocument struct {
	UID      uint   `json:"uid,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Username string `json:"username,omitempty"`

	EID      string    `json:"eid,omitempty"`
	Type     int8      `json:"type,omitempty"`
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	Datetime time.Time `json:"datetime"`
}

// UpdateEventDocument 更新消息文档体
type UpdateEventDocument struct {
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	Tags       []string    `json:"tags"`
	Constraint interface{} `json:"constraint,omitempty"`
}

type FileField struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// NoticeField 通知消息字段
type NoticeField struct {
	Files []*FileField `json:"files"`
}

// VoteField 投票消息字段
type VoteField struct {
	// 最多的投票量
	MaxNumber int      `json:"max_number"`
	Deadline  int64    `json:"deadline"`
	Options   []string `json:"options"`
}

// SortitionField 随机消息字段
type SortitionField struct {
	RequiredNumber int   `json:"required_number"`
	AllowedCancel  bool  `json:"allowed_cancel"`
	Deadline       int64 `json:"deadline"`
}

// ParticipationField 报名消息字段
type ParticipationField struct {
	RequiredNumber int   `json:"required_number"`
	AllowedCancel  bool  `json:"allowed_cancel"`
	Deadline       int64 `json:"deadline"`
}

// Insert 插入消息
func (e *EventDocument) Insert() error {
	return client.Insert(eventCollectionName, e)
}

// Update 更新消息
func (e *EventDocument) Update(eid string, value *UpdateEventDocument) (int64, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}
	var res map[string]interface{}
	if err := json.Unmarshal(data, &res); err != nil {
		return 0, err
	}

	return client.Update(
		eventCollectionName,
		bson.D{{"eid", eid}},
		bson.D{{"$set", res}},
	)
}

// Delete 删除消息
func (e *EventDocument) Delete(eid string) error {
	return client.Delete(eventCollectionName, bson.D{{"eid", eid}})
}

// Find 分页获取最近消息
func (e *EventDocument) Find(uid uint, offset, number int64) ([]SimpleEventDocument, error) {
	option := options.Find().SetSort(map[string]interface{}{"_id": -1}).SetSkip(offset).SetLimit(number)
	cursor, err := client.Find(eventCollectionName, bson.D{{"uid", uid}}, option)
	if err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(client.context, client.timeout*3)
	defer cancel1()
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx1)

	ctx2, cancel2 := context.WithTimeout(client.context, client.timeout*2)
	defer cancel2()
	var res []SimpleEventDocument
	for cursor.Next(ctx2) {
		var data SimpleEventDocument
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}

// FindOneDetail 获取单条消息详情
func (e *EventDocument) FindOneDetail(eid string) (event EventDocument, err error) {
	res := client.FindOne(eventCollectionName, bson.D{{"eid", eid}})
	if err = res.Err(); err != nil {
		return
	}

	err = res.Decode(&event)
	return
}

func (e *EventDocument) FindIn(eids []string) ([]SimpleEventDocument, error) {
	option := options.Find().SetSort(map[string]interface{}{"_id": -1})
	filter := bson.D{{"eid", bson.D{{"$in", bson.A{eids}}}}}

	cursor, err := client.Find(eventCollectionName, filter, option)
	if err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(client.context, client.timeout*3)
	defer cancel1()
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			return
		}
	}(cursor, ctx1)

	ctx2, cancel2 := context.WithTimeout(client.context, client.timeout*2)
	defer cancel2()
	var res []SimpleEventDocument
	for cursor.Next(ctx2) {
		var data SimpleEventDocument
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}
