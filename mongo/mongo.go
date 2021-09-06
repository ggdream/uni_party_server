package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *MongoDB
)

// MongoDB 封装的MongoDB
type MongoDB struct {
	client  *mongo.Client
	context context.Context

	timeout  time.Duration
	database string
}

// New 实例化自己封装的 MongoDB
func New(mongo *mongo.Client) *MongoDB {
	return &MongoDB{
		client:   mongo,
		context:  context.Background(),
		timeout:  300 * time.Millisecond,
		database: "test",
	}
}

// Init 初始化这个包
func Init() error {
	clientOptions := options.Client().ApplyURI("mongodb://root:hpxhJvhnrIAbX6AWd9DVLhcaZAoaGD8m@101.34.2.166:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	if err = c.Ping(ctx, nil); err != nil {
		return err
	}

	client = New(c)
	return nil
}

// Insert 向集合里插入文档
func (m *MongoDB) Insert(collection string, document interface{}) error {
	ctx, cancel := context.WithTimeout(m.context, m.timeout)
	defer cancel()
	_, err := m.client.Database(m.database).Collection(collection).InsertOne(ctx, document)
	return err
}

// Update 更新集合里的文档
func (m *MongoDB) Update(collection string, filter, update interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(m.context, m.timeout)
	defer cancel()
	res, err := m.client.Database(m.database).Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, err
}

// Delete 更新删除集合里的文档
func (m *MongoDB) Delete(collection string, filter interface{}) error {
	ctx, cancel := context.WithTimeout(m.context, m.timeout)
	defer cancel()
	_, err := m.client.Database(m.database).Collection(collection).DeleteOne(ctx, filter)
	return err
}

// DeleteMany 批量删除集合里的文档
func (m *MongoDB) DeleteMany(collection string, filter interface{}) error {
	ctx, cancel := context.WithTimeout(m.context, m.timeout)
	defer cancel()
	_, err := m.client.Database(m.database).Collection(collection).DeleteMany(ctx, filter)
	return err
}

// Find 查找集合里的多个文档
func (m *MongoDB) Find(collection string, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	ctx, cancel := context.WithTimeout(m.context, m.timeout*2)
	defer cancel()
	return m.client.Database(m.database).Collection(collection).Find(ctx, filter, opts...)
}

// FindOne 查找集合里的单个文档
func (m *MongoDB) FindOne(collection string, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(m.context, m.timeout*2)
	defer cancel()
	return m.client.Database(m.database).Collection(collection).FindOne(ctx, filter, opts...)
}

// Count 统计记录数
func (m *MongoDB) Count(collection string, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	ctx, cancel := context.WithTimeout(m.context, m.timeout*2)
	defer cancel()
	return m.client.Database(m.database).Collection(collection).CountDocuments(ctx, filter, opts...)
}
