package cache

import (
	"context"
	"fmt"
)

type Student struct{}

func (s Student) joinClassKey(uid uint) string {
	return fmt.Sprintf("student:class:%d", uid)
}

// Set 为一个高校存储学生与班级的对应信息
// key为班级号uid，value为学生uid
func (s Student) Set(uidMapList map[uint][]string) error {
	pipe := client.client.Pipeline()
	defer pipe.Close()

	for key, value := range uidMapList {
		_ = pipe.SAdd(client.context, s.joinClassKey(key), value).Err()
	}

	ctx, cancel := context.WithTimeout(client.context, client.timeout*3)
	defer cancel()
	_, err := pipe.Exec(ctx)
	return err
}

// Get 获取某个班级内的所有学生的uid
func (s Student) Get(classUid uint) ([]string, error) {
	return client.SMembers(s.joinClassKey(classUid))
}
