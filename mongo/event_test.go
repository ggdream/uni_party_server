package mongo

import (
	"fmt"
	"testing"
	"time"
)

// func init() {
// 	if err := Init(); err != nil {
// 		panic(err)
// 	}
// }

// 插入测试
func TestEventCollection_Insert(t *testing.T) {
	event := &EventDocument{
		UID:      118,
		EID:      "aadf",
		Type:     0,
		Title:    "网址",
		Content:  "分享一下",
		Datetime: time.Now(),
	}
	if err := event.Insert(); err != nil {
		panic(err)
	}
}

func TestEventDocument_Update(t *testing.T) {
	event := EventDocument{}
	_, err := event.Update("aadf", &UpdateEventDocument{
		Title:      "hello",
		Content:    "a",
		Tags:       []string{"1", "2", "3"},
		Constraint: nil,
	})
	if err != nil {
		panic(err)
	}
}

// 找到一个消息的详情
func TestEventDocument_FindOneDetail(t *testing.T) {
	event := &EventDocument{
		UID:        111,
		EID:        "faf",
		Type:       0,
		Title:      "哈喽呀",
		Content:    "你这",
		Constraint: "",
		Datetime:   time.Now(),
	}
	res, err := event.FindOneDetail("faf")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

// 分页查询
func TestEventDocument_Find(t *testing.T) {
	event := &EventDocument{
		UID:        111,
		EID:        "faf",
		Type:       0,
		Title:      "哈喽呀",
		Content:    "你这",
		Constraint: "",
		Datetime:   time.Now(),
	}
	res, err := event.Find(111, 0, 2)
	if err != nil {
		panic(err)
	}
	for _, v := range res {
		fmt.Println(v)
	}
}
