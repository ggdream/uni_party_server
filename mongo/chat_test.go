package mongo

import (
	"testing"
	"time"
)

// func init() {
// 	if err := Init(); err != nil {
// 		panic(err)
// 	}
// }

// 插入测试
func TestChatDocument_Insert(t *testing.T) {
	chat := &ChatDocument{
		Type:     0,
		FromUID:  10,
		ToUID:    205,
		GroupID:  "sfs",
		Message:  "在干嘛呢",
		URL:      "",
		Others:   nil,
		MID:      "xxxx",
		Datetime: time.Now().UnixNano(),
		Version:  "1",
	}
	err := chat.Insert()
	if err != nil {
		panic(err)
	}
}
