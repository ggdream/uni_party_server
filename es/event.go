package es

import (
	"context"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

const eventIndexName = "event"

type EventIndex struct {
	UID    uint   `json:"uid"`
	Avatar string `json:"avatar"`

	EID      string `json:"eid"`
	Type     int8   `json:"type"`
	Title    string `json:"title"`
	Datetime int64  `json:"datetime"`
}

// Insert 插入新消息
func (e *EventIndex) Insert() error {
	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()
	_, err := client.es.Index().Index(eventIndexName).BodyJson(e).Do(ctx)
	return err
}

// Update 更新消息标题
func (e *EventIndex) Update(eid uint, title string) error {
	query := elastic.NewTermQuery("eid", eid)
	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	script := elastic.NewScript(fmt.Sprintf("ctx._source.title='%s'", title))
	_, err := client.es.UpdateByQuery(userIndexName).Query(query).Script(script).Do(ctx)
	return err
}

// Delete 根据eid删除记录
func (e *EventIndex) Delete(eid string) error {
	query := elastic.NewTermQuery("eid", eid)

	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	_, err := client.es.DeleteByQuery().Index(eventIndexName).Query(query).Do(ctx)
	return err
}

// QueryByTitle 分页按消息标题模糊匹配查询消息
func (e *EventIndex) QueryByTitle(uid uint, title string, offset, number int) (int, []EventIndex, error) {
	query1 := elastic.NewMatchQuery("title", title)
	query2 := elastic.NewTermQuery("uid", uid)

	aggs := elastic.NewCardinalityAggregation().Field("eid")
	ctx1, cancel1 := context.WithTimeout(client.context, client.timeout)
	defer cancel1()
	cData, err := client.es.Search().Index(eventIndexName).Query(query1).Query(query2).Aggregation("total", aggs).Size(0).Do(ctx1)
	if err != nil {
		return 0, nil, err
	}
	agg, found := cData.Aggregations.ValueCount("total")
	if !found {
		return 0, nil, errors.New("count not found")
	}

	ctx2, cancel2 := context.WithTimeout(client.context, client.timeout)
	defer cancel2()
	data, err := client.es.Search().Index(eventIndexName).Query(query1).Query(query2).From(offset).Size(number).Pretty(true).Do(ctx2)
	if err != nil {
		return 0, nil, err
	}

	var result []EventIndex
	for _, record := range data.Each(reflect.TypeOf(EventIndex{})) {
		result = append(result, record.(EventIndex))
	}

	return int(*agg.Value), result, nil
}

// QueryByTitleAndType 分页按消息标题和消息类型模糊匹配查询消息
func (e *EventIndex) QueryByTitleAndType(title string, eventType int8, offset, number int) ([]EventIndex, error) {
	query1 := elastic.NewMatchQuery("title", title)
	query2 := elastic.NewTermQuery("type", eventType)

	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	data, err := client.es.Search().Index(eventIndexName).Query(query1).Query(query2).From(offset).Size(number).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	var result []EventIndex
	for _, record := range data.Each(reflect.TypeOf(EventIndex{})) {
		result = append(result, record.(EventIndex))
	}

	return result, nil
}
