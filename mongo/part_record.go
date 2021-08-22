package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	partCollectionName = "part_record"
)

type PartDocument struct {
	UID uint   `json:"uid"`
	EID string `json:"eid"`
}

// Insert 插入用户新报名
func (s *PartDocument) Insert(requiredNumber int) error {
	session, err := client.client.StartSession()
	if err != nil {
		return err
	}

	if err = session.StartTransaction(); err != nil {
		return err
	}

	transaction := func(sc mongo.SessionContext) error {
		filter := bson.D{{Key: "eid", Value: s.EID}}
		count, err := client.client.Database(client.database).Collection(partCollectionName).CountDocuments(sc, filter)
		if err != nil {
			return err
		}
		if int(count) == requiredNumber {
			return errors.New("it is full")
		}

		_, err = client.client.Database(client.database).Collection(partCollectionName).InsertOne(sc, s)
		if err != nil {
			return err
		}

		return session.CommitTransaction(sc)
	}

	ctx, cancel := context.WithTimeout(client.context, client.timeout*2)
	defer cancel()

	err = mongo.WithSession(ctx, session, transaction)
	defer session.EndSession(ctx)

	return err
}

// Query 查询报名参加情况
func (s *PartDocument) Query(eid string, offset, number int64) (res []VoteDocument, err error) {
	option := options.Find().SetSkip(offset).SetLimit(number)
	cursor, err := client.Find(partCollectionName, bson.D{{Key: "eid", Value: eid}}, option)
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
	//for cursor.Next(ctx2) {
	//	var data VoteDocument
	//	if err := cursor.Decode(&data); err != nil {
	//		return nil, err
	//	}
	//	res = append(res, data)
	//}
	err = cursor.All(ctx2, &res)
	return
}
