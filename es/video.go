package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"reflect"
)

const videoIndexName = "video"

type VideoIndex struct {
	UID    uint   `json:"uid"`
	Avatar string `json:"avatar"`

	VID   string   `json:"vid"`
	Title string   `json:"title"`
	Tags  []string `json:"tags"`

	Cover string `json:"cover"`
	Video string `json:"video"`

	Datetime int64 `json:"datetime"`
}

// Insert 插入新消息
func (v *VideoIndex) Insert() error {
	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()
	_, err := client.es.Index().Index(videoIndexName).BodyJson(v).Do(ctx)
	return err
}

// Delete 根据eid删除记录
func (v *VideoIndex) Delete(vid string) error {
	query := elastic.NewTermQuery("vid", vid)

	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	_, err := client.es.DeleteByQuery().Index(videoIndexName).Query(query).Do(ctx)
	return err
}

// QueryByTitle 分页按视频标题模糊匹配查询视频
func (v *VideoIndex) QueryByTitle(title string, offset, number int) ([]VideoIndex, error) {
	query := elastic.NewMatchQuery("title", title)

	ctx, cancel := context.WithTimeout(client.context, client.timeout)
	defer cancel()

	data, err := client.es.Search().Index(videoIndexName).Query(query).From(offset).Size(number).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}

	var result []VideoIndex
	for _, record := range data.Each(reflect.TypeOf(VideoIndex{})) {
		result = append(result, record.(VideoIndex))
	}

	return result, nil
}
