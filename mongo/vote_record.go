package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	voteCollectionName = "vote_record"
)

type VoteDocument struct {
	UID      uint      `json:"uid"`
	EID      string    `json:"eid"`
	Answers  []bool    `json:"answers"`
	Datetime time.Time `json:"datetime"`
}

// Insert 插入用户新投票
func (v *VoteDocument) Insert() error {
	return client.Insert(voteCollectionName, v)
}

// Query 查询投票情况
func (v *VoteDocument) Query(eid string, offset, number int64) ([]VoteDocument, error) {
	option := options.Find().SetSkip(offset).SetLimit(number)
	cursor, err := client.Find(voteCollectionName, bson.D{{"eid", eid}}, option)
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
	var res []VoteDocument
	for cursor.Next(ctx2) {
		var data VoteDocument
		if err := cursor.Decode(&data); err != nil {
			return nil, err
		}
		res = append(res, data)
	}

	return res, nil
}
