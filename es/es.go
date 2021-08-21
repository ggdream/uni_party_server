package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"time"
)

var (
	client *ElasticSearch
)

// ElasticSearch 自己的封装体
type ElasticSearch struct {
	es      *elastic.Client
	context context.Context
	timeout time.Duration
}

// Init 实例化本包
func Init() (err error) {
	client, err = New()
	return
}

// New 实例化我自己的ES Client
func New() (*ElasticSearch, error) {
	es, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		return nil, err
	}
	_, _, err = es.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		return nil, err
	}

	return &ElasticSearch{
		es:      es,
		context: context.Background(),
		timeout: 300 * time.Millisecond,
	}, nil
}

// Insert 插入数据
func (e *ElasticSearch) Insert(index string, data interface{}) (*elastic.IndexResponse, error) {
	ctx, cancel := context.WithTimeout(e.context, e.timeout)
	defer cancel()
	return client.es.Index().Index(index).BodyJson(data).Do(ctx)
}




















