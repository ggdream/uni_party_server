package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	chatCollectionName = "chat"
)

// ChatDocument 通信请求体
type ChatDocument struct {
	// Type 消息类型
	Type int8

	// FromUID 发送消息的用户的UID
	FromUID int
	// ToUID 接受消息的用户的UID
	ToUID int
	// GroupID 组聊ID
	GroupID string

	// Message 文本消息
	Message string
	// URL 二进制文件地址
	URL string
	// Others 扩展字段，附加的其他数据
	Others interface{}

	// MID 消息ID，服务器收到消息时，无此字段；转发消息时，需要添加此字段
	MID string
	// Datetime 消息接受时间，服务器收到消息时，无此字段；转发消息时，需要添加此字段
	Datetime int64
	// Version 为方便后期协议修改，添加版本号字段
	Version string
}


// Insert 插入消息
func (c *ChatDocument) Insert() error {
	return client.Insert(chatCollectionName, c)
}


// Delete 删除消息
func (c *ChatDocument) Delete(mid string) error {
	return client.Delete(chatCollectionName, bson.D{{"mid", mid}})
}

// Find 分页获取最近消息
func (c *ChatDocument) Find(uid uint, offset, number int64) ([]ChatDocument, error) {
	option := options.Find().SetSkip(offset).SetLimit(number).SetSort(map[string]interface{}{"_id": -1})
	cursor, err := client.Find(chatCollectionName, bson.D{{"uid", uid}}, option)
	if err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(client.context, client.timeout*3)
	defer cancel1()
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			// TODO: log
		}
	}(cursor, ctx1)

	ctx2, cancel2 := context.WithTimeout(client.context, client.timeout*2)
	defer cancel2()
	var res []ChatDocument
	for cursor.Next(ctx2) {
		var data ChatDocument
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}
