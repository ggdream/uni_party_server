package cache

import (
	"strconv"
)

// UniversityRecord 高校信息记录
// 使用 Hash，field为uid，value为大学信息JSON字符串
type UniversityRecord struct{}

// joinUniversityRecordKey 高校信息记录的string-key
const joinUniversityRecordKey = "university:record"

// Add 高校入驻
func (r UniversityRecord) Add(uid uint, value string) error {
	return client.HSet(joinUniversityRecordKey, strconv.Itoa(int(uid)), value)
}

// Del 高校离去
func (r UniversityRecord) Del(uid uint) error {
	return client.HDel(joinUniversityRecordKey, strconv.Itoa(int(uid)))
}

// Get 获取高校信息
func (r UniversityRecord) Get(uid uint) (string, error) {
	return client.HGet(joinUniversityRecordKey, strconv.Itoa(int(uid))).Result()
}
