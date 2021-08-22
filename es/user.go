package es

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

const userIndexName = "user"

type UserIndex struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	College  string `json:"college"`
}

// Insert 插入新用户
func (u *UserIndex) Insert() error {
	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()
	_, err := client.es.Index().Index(userIndexName).BodyJson(u).Do(ctx)
	return err
}

// Update 更新用户昵称
func (u *UserIndex) Update(uid uint, username string) error {
	query := elastic.NewTermQuery("uid", uid)
	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	script := elastic.NewScript(fmt.Sprintf("ctx._source.username='%s'", username))
	_, err := client.es.UpdateByQuery(userIndexName).Query(query).Script(script).Do(ctx)
	return err
}

// Query 分页按用户昵称模糊匹配查询用户
func (u *UserIndex) Query(username string, offset, number int) (int, []UserIndex, error) {
	query := elastic.NewMatchQuery("username", username)

	aggs := elastic.NewCardinalityAggregation().Field("uid")
	ctx1, cancel1 := context.WithTimeout(client.context, client.timeout)
	defer cancel1()
	cData, err := client.es.Search().Index(eventIndexName).Query(query).Aggregation("total", aggs).Size(0).Do(ctx1)
	if err != nil {
		return 0, nil, err
	}
	agg, found := cData.Aggregations.ValueCount("total")
	if !found {
		return 0, nil, errors.New("count not found")
	}

	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	data, err := client.es.Search().Index(userIndexName).Query(query).From(offset).Size(number).Pretty(true).Do(ctx)
	if err != nil {
		return 0, nil, err
	}

	var result []UserIndex
	for _, record := range data.Each(reflect.TypeOf(UserIndex{})) {
		result = append(result, record.(UserIndex))
	}

	return int(*agg.Value), result, nil
}
