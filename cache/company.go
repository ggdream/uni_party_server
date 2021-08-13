package cache

import (
	"strconv"
)

// CompanyRecord 企业信息记录
// 使用 Hash，field为uid，value为大学信息JSON字符串
type CompanyRecord struct{}

// joinCompanyRecordKey 企业信息记录的string-key
const joinCompanyRecordKey = "company:record"

// Add 企业入驻
func (r CompanyRecord) Add(uid uint, value string) error {
	return client.HSet(joinCompanyRecordKey, strconv.Itoa(int(uid)), value)
}

// Del 企业离去
func (r CompanyRecord) Del(uid uint) error {
	return client.HDel(joinCompanyRecordKey, strconv.Itoa(int(uid)))
}

// Get 获取企业信息
func (r CompanyRecord) Get(uid uint) (string, error) {
	return client.HGet(joinCompanyRecordKey, strconv.Itoa(int(uid))).Result()
}
